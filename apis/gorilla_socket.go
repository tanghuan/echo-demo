package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func read(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()

		fmt.Println("msgType = ", msgType)
		fmt.Println("msg = ", string(msg))

		if err != nil {
			fmt.Println("read error: ", err)
			conn.Close()
			return
		}

		fmt.Printf("msg = %s \n", msg)

		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func write(conn *websocket.Conn) {
	ch := time.Tick(20 * time.Second)
	for t := range ch {

		formatTime := t.Format("2006-01-02 15:04:05")
		fmt.Println("time = ", formatTime)

		err := conn.WriteMessage(websocket.TextMessage, []byte("SendAt: "+formatTime))
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

func writePing(conn *websocket.Conn) {
	ch := time.Tick(1 * time.Second)
	for t := range ch {

		formatTime := t.Format("2006-01-02 15:04:05")
		fmt.Println("time = ", formatTime)
		err := conn.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

// HandleConnect 处理Socket 链接
func HandleConnect(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	go read(conn)

	// go write(conn)

	go writePing(conn)

	return nil
}
