package websocket


//when any method/function needs to accessed globally,
//make its first letter Capitalised

import(
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize = 1024,
	WriteBufferSize = 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		log.println(err)
		return ws,err
	}
	return ws, nil
}

func Reader(conn *websocket.conn){
	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.println(err)
			return
		}

		fmt.println(string(p))

		if err := conn.writeMessage(messageType, p); err != nil{
			log.println(err)
			return
		}
	}
}

func Writer(conn *websocket.conn){
	for{
		fmt.Println("sending")
		messageType, r, err := conn.NextReader()
		if err != nil{
			fmt.Println(err)
			return
		}
		w,err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := w.close(); err != nil{
			fmt.Println(err)
			return
		}
	}
}


