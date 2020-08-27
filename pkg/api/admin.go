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
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xcloud-webconsole/pkg/dao"
)

//stat --- connections,mem,cpu,
//manager--- maxConnections,

func Add(c *gin.Context){
	name := c.PostForm("name")
	address := c.PostForm("address")
	username := c.PostForm("username")
	password := c.PostForm("password")
	sshKey := c.PostForm("sshKey")

	session := new(dao.Session)
	session.Name = name
	session.Address = address
	session.Username = username
	session.Password = password
	session.SshKey = sshKey
	id := dao.InsertSession(session)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id": id,
	})
}

func List(c *gin.Context){

	sessions :=dao.SessionList()
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"sessions" : sessions,
	})
}

func Del(c *gin.Context){
	idStr := c.PostForm("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	dao.DelSession(id)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}



