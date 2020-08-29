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
package main

import (
	"log"
	"net/http"
	"xcloud-webconsole/pkg/api"
	"xcloud-webconsole/pkg/core"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("WebConsole starting...")

	engine := gin.Default()
	engine.Use(createCorsHandler())

	engine.GET("/ws/:id", core.WsSsh)
	engine.POST("/admin/add", api.Add)
	engine.POST("/admin/del", api.Del)
	engine.GET("/admin/list", api.List)

	err := engine.Run(":8888") // Default listen on 0.0.0.0:8080.
	if err != nil {
		log.Panic(err)
	}

}

// createCorsHandler ...
func createCorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// var headerNames []string
		// for k := range c.Request.Header {
		// 	headerNames = append(headerNames, k)
		// }
		// headerNamesString := strings.Join(headerNames, ", ")
		// if headerNamesString != "" {
		// 	headerNamesString = fmt.Sprintf("Access-Control-Allow-Origin, Access-Control-Allow-Headers, %s", headerNamesString)
		// } else {
		// 	headerNamesString = "Access-Control-Allow-Origin, Access-Control-Allow-Headers"
		// }

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Cors OPTIONS Request")
		}

		// Sets default access control policy for CORS requests
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "")
			c.Header("Access-Control-Expose-Headers", "FooBar")   // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		// Execute the next handler
		c.Next()
	}
}
