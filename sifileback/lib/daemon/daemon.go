package daemon

import (
	"./r7uter"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	ServerPort int
	StaticFolder string
	listenSpec string
	sssets http.FileSystem
	Close chan os.Signal
	Loger chan <- [4]string
}



func Run(cfg *Config) (destroyed chan bool, err error) {
	// Начинаем открывать сервер
	cfg.listenSpec = fmt.Sprintf(":%d",cfg.ServerPort)

	cfg.Loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Starting, HTTP on%s\n", cfg.listenSpec)}
	// Настраиваем сервер
	l, err := net.Listen("tcp", cfg.listenSpec)
	if err != nil {
		cfg.Loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Error creating listener: %v\n", err),"1"}
		return
	}

	start(cfg, l)
	waitForSignal(cfg.Close,cfg.Loger,destroyed)
	return
}

func waitForSignal(close chan os.Signal, loger chan<- [4]string, destroyed chan bool) {
	subNmae := "Daemon_Closer"
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)
	s := <-close
	//fmt.Printf("%s | Безопасное завершение программы\n",subNmae)
	// Диплога могло вообще не создасться
	// !!! Мы не проверяем записаны ли все данные, просто закрываем это дело
	// loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Got signal: %v, exiting.", s)}
	// Пока программа при обрыве завершается через этот блок
	loger <- [4]string{"Front Http Server", "nil", fmt.Sprintf("%s | COM:Безопасное завершение программы | DT:%s", subNmae, s)}
	loger <- [4]string{"Front Http Server", "nil", fmt.Sprintf("%s | COM:Завершаем сессии", s)}
	// fmt.Printf("%d | Закрываем базу с пользователями\n",operInc())

	//ReactBro.Close()
	destroyed <- true
}
func start(cfg *Config,l net.Listener) {
	demon := mux.NewRouter()
	//Шэринг ресурсов в разных доменах
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization","Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Set-Cookie"})
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080","http://localhost:8081","*"}) 	// * для куков не катит
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedCredentials := handlers.AllowCredentials() //для куков
	var store = sessions.NewCookieStore([]byte{4,145,118,216,124,175,115,175,203,125,230,103,195,150,220,153,44,253,31,82,123,171,100,60,20,156,49,181,32,70,242,251})
	store.Options = &sessions.Options{
		SameSite: http.SameSiteNoneMode,
		Secure: false, // Если установлен true, то в режиме разработчика ничего не видно, а иначе не нельзя кроссдоменные куки устанавливать
		Path:   "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// securecookie.GenerateRandomKey(32) // [4 145 118 216 124 175 115 175 203 125 230 103 195 150 220 153 44 253 31 82 123 171 100 60 20 156 49 181 32 70 242 251]
	server := &http.Server{
		Handler:       handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods,allowedCredentials)(demon),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16 }
	r7 := r7uter.NewRouterManager(cfg.Loger)
	demon.HandleFunc("/ngapi/run", r7.RunCmd).Methods("POST")
	demon.HandleFunc("/echo", echo)
	//router.PathPrefix("").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./resourse/"))))

	go server.Serve(l)
	cfg.Loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Server is started")}
}


var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
