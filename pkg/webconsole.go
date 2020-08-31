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
	"context"
	"net/http"
	"strconv"
	"time"
	admin "xcloud-webconsole/pkg/admin"
	config "xcloud-webconsole/pkg/config"
	logging "xcloud-webconsole/pkg/logging"
	ssh2 "xcloud-webconsole/pkg/modules/ssh2"
	"xcloud-webconsole/pkg/utils"

	"go.uber.org/zap"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// WebConsole ...
type WebConsole struct {
	stopper *utils.Stopper
}

// StartServe ...
func (wc *WebConsole) StartServe(ctx context.Context, conf string) {
	wc.stopper = utils.NewStopper(ctx, func() {
		logging.Main.Info("Stopping ...")
	})

	// Init global config.
	config.InitGlobalConfig(conf)

	// Init zap logger.
	logging.InitZapLogger()

	// Start webserver...
	go wc.startWebServer()

	// Start admin server
	go admin.ServeStart()

	// Waiting for system exit
	wc.stopper.WaitForExit()
}

// startWebServer ...
func (wc *WebConsole) startWebServer() *gin.Engine {
	logging.Main.Info("WebConsole server starting...")

	engine := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	engine.Use(wc.createCorsHandler())
	zapLogger := logging.Main.GetZapLogger()
	engine.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zapLogger, true))

	// Register SSH2 handlers
	wc.registerSSH2Handlers(engine)

	err := engine.Run(config.GlobalConfig.Server.Listen) // Default listen on 0.0.0.0:8080.
	if err != nil {
		logging.Receive.Panic("error", zap.Error(err))
	}

	return engine
}

// registerSSH2Handlers ...
func (wc *WebConsole) registerSSH2Handlers(engine *gin.Engine) {
	engine.GET(ssh2.DefaultSSH2APIWebSocketURI, ssh2.NewWebsocketConnectionFunc)
	engine.GET(ssh2.DefaultSSH2APISessionQueryURI, ssh2.QuerySSH2SessionsFunc)
	engine.POST(ssh2.DefaultSSH2APISessionAddURI, ssh2.AddSSH2SessionFunc)
	engine.POST(ssh2.DefaultSSH2APISessionDeleteURI, ssh2.DeleteSSH2SessionFunc)
}

// createCorsHandler ...
func (wc *WebConsole) createCorsHandler() gin.HandlerFunc {
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
			c.Writer.Header().Set("Access-Control-Allow-Origin", config.GlobalConfig.Server.Cors.AllowOrigins)
			c.Header("Access-Control-Allow-Origin", config.GlobalConfig.Server.Cors.AllowOrigins)
			c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(config.GlobalConfig.Server.Cors.AllowCredentials))
			c.Header("Access-Control-Allow-Methods", config.GlobalConfig.Server.Cors.AllowMethods)
			c.Header("Access-Control-Allow-Headers", config.GlobalConfig.Server.Cors.AllowHeaders)
			c.Header("Access-Control-Expose-Headers", config.GlobalConfig.Server.Cors.ExposeHeaders) // 跨域关键设置让浏览器可以解析
			c.Header("Access-Control-Max-Age", strconv.Itoa(config.GlobalConfig.Server.Cors.MaxAge))
			c.Set("Content-Type", "application/json")
		}

		// Execute the next handler
		c.Next()
	}
}
