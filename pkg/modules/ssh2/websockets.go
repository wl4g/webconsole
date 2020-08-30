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
package ssh2

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// NewWebsocketConnectionFunc ...
func NewWebsocketConnectionFunc(c *gin.Context) {
	wsSecID := c.Request.Header.Get("Sec-WebSocket-Key")
	webssh := NewWebSSH2Dispatcher()

	webssh.SetTerm(DefaultTerm)
	webssh.SetBuffSize(8192)
	webssh.SetWSSecID(wsSecID)
	webssh.SetConnTimeOut(5 * time.Second)
	webssh.SetLogger(log.New(os.Stderr, "[webssh] ", log.Ltime|log.Ldate))

	upgrader := websocket.Upgrader{
		// Cross/cors origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Resolve: Sec-WebSocket-Protocol Header
		//Subprotocols: []string{r.Header.Get("Sec-WebSocket-Protocol")},
		Subprotocols:    []string{"webssh"},
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Panic(err)
	}

	webssh.AddWebsocket(ws)
}
