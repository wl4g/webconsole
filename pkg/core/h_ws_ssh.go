/**
 * Copyright 2017 ~ 2025 the original author or authors<Wanglsir@gmail.com, 983708408@qq.com>.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package core

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsSsh ...
func WsSsh(c *gin.Context) {
	id := c.Request.Header.Get("Sec-WebSocket-Key")
	webssh := NewWebSSH()
	// term 可以使用 ansi, linux, vt100, xterm, dumb，除了 dumb外其他都有颜色显示, 默认 xterm
	webssh.SetTerm(TermXterm)
	webssh.SetBuffSize(8192)
	webssh.SetId(id)
	webssh.SetConnTimeOut(5 * time.Second)
	webssh.SetLogger(log.New(os.Stderr, "[webssh] ", log.Ltime|log.Ldate))

	wsUpgrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Resolve: Sec-WebSocket-Protocol Header
		//Subprotocols: []string{r.Header.Get("Sec-WebSocket-Protocol")},
		Subprotocols:    []string{"webssh"},
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}

	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Panic(err)
	}

	webssh.AddWebsocket(ws)
}
