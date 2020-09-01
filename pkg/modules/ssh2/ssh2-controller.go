/**
 * Copyright 2017 ~ 2025 the original author or authors.
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
	"net/http"
	"strconv"
	store "xcloud-webconsole/pkg/modules/ssh2/store"

	"github.com/gin-gonic/gin"
)

// AddSSH2SessionFunc ...
func AddSSH2SessionFunc(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	username := c.PostForm("username")
	password := c.PostForm("password")
	sshKey := c.PostForm("sshKey")

	session := new(store.SessionBean)
	session.Name = name
	session.Address = address
	session.Username = username
	session.Password = password
	session.SSHPrivateKey = sshKey
	id := store.GetDelegate().SaveSession(session)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}

// QuerySSH2SessionsFunc ...
func QuerySSH2SessionsFunc(c *gin.Context) {
	sessions := store.GetDelegate().QuerySessionList()
	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"sessions": sessions,
	})
}

// DeleteSSH2SessionFunc ...
func DeleteSSH2SessionFunc(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	store.GetDelegate().DeleteSession(id)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

// CloseSSH2SessionFunc ...
func CloseSSH2SessionFunc(c *gin.Context) {
	// TODO Closing dispatcher channel
	// ...

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

const (
	// DefaultSSH2APIBaseURI ssh2 console base URI.
	DefaultSSH2APIBaseURI = "/terminal/"

	// DefaultSSH2APIWebSocketURI ssh2 console websocket connection URI.
	DefaultSSH2APIWebSocketURI = DefaultSSH2APIBaseURI + "ws/:id"

	// DefaultSSH2APISessionQueryURI ssh2 console session query URI.
	DefaultSSH2APISessionQueryURI = DefaultSSH2APIBaseURI + "session/list"

	// DefaultSSH2APISessionAddURI ssh2 console session create URI.
	DefaultSSH2APISessionAddURI = DefaultSSH2APIBaseURI + "session/create"

	// DefaultSSH2APISessionDeleteURI ssh2 console session delete URI.
	DefaultSSH2APISessionDeleteURI = DefaultSSH2APIBaseURI + "session/delete"

	// DefaultSSH2APISessionCloseURI ssh2 console session close URI.
	DefaultSSH2APISessionCloseURI = DefaultSSH2APIBaseURI + "session/close"
)
