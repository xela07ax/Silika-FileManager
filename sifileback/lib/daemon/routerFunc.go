package daemon

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/processor"
	"github.com/xela07ax/toolsXela/ttp"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{} // use default options

var minions int
type Router struct {
	name string
	Tx chan processor.ToDocker
	loger chan <- [4]string
}
func NewRouterManager(loger chan <- [4]string, tx chan processor.ToDocker) *Router {
	return &Router{
		name: 		"RouteManager",
		loger: 		loger,
		Tx: tx,
	}
}

const (
	daemon = "Daemon"
	conn = "Conn"
)


func (fs *Router) RunKod(w http.ResponseWriter, r *http.Request)  {
	fs.loger <- [4]string{daemon,conn,"Run code welcome"}
	wait := make(chan processor.ToDocker,1)
	fs.Tx <- processor.ToDocker{
		Data:  nil,
		Acton: "Run",
		Out: wait,
	}
	fs.loger <- [4]string{daemon,conn,"Run code wait"}
	fmt.Println(<- wait)
	fs.loger <- [4]string{daemon,conn,"Run code ok"}
}
func (fs *Router) Conn(w http.ResponseWriter, r *http.Request)  {
	minions++
	unit := minions
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		tx := fmt.Sprintf("upgrade: Не удалось создать соединение по WS|unit:%d | Ertx:%v",unit,err)
		ttp.Resp(w,r,conn, tx,1,true)
		fs.loger <- [4]string{daemon,conn,tx, "1"}
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