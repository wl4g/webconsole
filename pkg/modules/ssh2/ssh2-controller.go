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
	repo "xcloud-webconsole/pkg/modules/ssh2/repository"

	"github.com/gin-gonic/gin"
)

//stat --- connections,mem,cpu,
//manager--- maxConnections,

// AddSSH2SessionFunc ...
func AddSSH2SessionFunc(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	username := c.PostForm("username")
	password := c.PostForm("password")
	sshKey := c.PostForm("sshKey")

	session := new(repo.SessionBean)
	session.Name = name
	session.Address = address
	session.Username = username
	session.Password = password
	session.SSHPrivateKey = sshKey
	id := repo.GetDelegateSSH2Repository().SaveSession(session)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}

// QuerySSH2SessionsFunc ...
func QuerySSH2SessionsFunc(c *gin.Context) {
	sessions := repo.GetDelegateSSH2Repository().QuerySessionList()
	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"sessions": sessions,
	})
}

// DeleteSSH2SessionFunc ...
func DeleteSSH2SessionFunc(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	repo.GetDelegateSSH2Repository().DeleteSession(id)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
