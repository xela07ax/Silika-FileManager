package daemon

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/processor"
	"github.com/xela07ax/toolsXela/hub"
	"github.com/xela07ax/toolsXela/hub/blog"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	ServerPort int
	StaticFolder string
	listenSpec string
	sssets http.FileSystem
	Destroy <- chan struct{}
	Loger chan <- [4]string
	Tx chan processor.ToDocker
	Hub *hub.Hub
}



func Run(cfg *Config){
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

	return
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
	fs := NewRouterManager(cfg.Loger, cfg.Tx)
	demon.HandleFunc("/fsapi/runKod", fs.RunKod).Methods("POST")
	demon.HandleFunc("/echo", Echo)
	demon.HandleFunc("/", home)
	demon.HandleFunc("/terminal", blog.Home) // WsLogger
	demon.HandleFunc("/wsLog", cfg.Hub.ServeWs) // WsLogger+
	//router.PathPrefix("").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./resourse/"))))

	go server.Serve(l)
	cfg.Loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Server is started")}
}


func Echo(w http.ResponseWriter, r *http.Request) {
	minions++
	unit := minions
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
		fmt.Printf("__Unit__:%d|recv: %s\n", unit, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo") // Изменим {{.}} на адрес
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    console.log("load")
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        console.log("open:","{{.}}")
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
 		console.log("send:",input.value)
        if (!ws) {
	        console.log("oops:","WebSocket.CLOSED")
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
	console.log("close")
        if (!ws) {
	        console.log("oops:","WebSocket.CLOSED")
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))