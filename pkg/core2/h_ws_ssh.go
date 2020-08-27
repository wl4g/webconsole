package core2

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func WsSsh(c *gin.Context) {
	id := c.Request.Header.Get("Sec-WebSocket-Key")
	webssh := NewWebSSH()
	// term 可以使用 ansi, linux, vt100, xterm, dumb，除了 dumb外其他都有颜色显示, 默认 xterm
	webssh.SetTerm(TermXterm)
	webssh.SetBuffSize(8192)
	webssh.SetId(id)
	webssh.SetConnTimeOut(5 * time.Second)
	webssh.SetLogger(log.New(os.Stderr, "[webssh] ", log.Ltime|log.Ldate))

	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		//Subprotocols: []string{r.Header.Get("Sec-WebSocket-Protocol")},
		Subprotocols: []string{"webssh"},
		ReadBufferSize: 8192,
		WriteBufferSize: 8192,
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Panic(err)
	}

	webssh.AddWebsocket(ws)
}
