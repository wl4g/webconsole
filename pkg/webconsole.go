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
	"strings"
	"time"
	admin "xcloud-webconsole/pkg/admin"
	config "xcloud-webconsole/pkg/config"
	logging "xcloud-webconsole/pkg/logging"
	ssh2 "xcloud-webconsole/pkg/modules/ssh2"
	"xcloud-webconsole/pkg/modules/ssh2/store"
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
	wc.stopper = utils.NewDefault(ctx, func() {
		logging.Main.Info("Stopping ...")
		// Store closing
		if err1 := store.GetDelegate().Close(); err1 != nil {
			logging.Main.Error("Closing store resource failure", zap.Error(err1))
		}
		// TODO Closing others resources gracefully
		// ...
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

	// Create gin engine.
	engine := gin.New()

	// Sets gin runtim mode.
	// gin.SetMode(gin.ReleaseMode)

	// Sets gin http cors policy.
	corsConfig := config.GlobalConfig.Server.Cors
	corsHolder := utils.CorsHolder{
		AllowOrigins:     strings.Split(corsConfig.AllowOrigins, ","),
		AllowMethods:     strings.Split(corsConfig.AllowMethods, ","),
		AllowHeaders:     strings.Split(corsConfig.AllowHeaders, ","),
		AllowCredentials: corsConfig.AllowCredentials,
		ExposeHeaders:    strings.Split(corsConfig.ExposeHeaders, ","),
		MaxAge:           corsConfig.MaxAge,
	}
	corsHolder.RegisterCorsProcessor(engine)

	// Sets gin http other configuration.
	engine.Use(func(c *gin.Context) {
		c.Set("Content-Type", "application/json")
	})

	// Sets gin http logger.
	zapLogger := logging.Main.GetZapLogger()
	engine.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zapLogger, true))

	// Sets gin http handlers
	wc.registerHTTPHandlers(engine)

	err := engine.Run(config.GlobalConfig.Server.Listen) // Default listen on 0.0.0.0:8080.
	if err != nil {
		logging.Receive.Panic("error", zap.Error(err))
	}

	return engine
}

// registerHTTPHandlers ...
func (wc *WebConsole) registerHTTPHandlers(engine *gin.Engine) {
	// Register SSH2 dispatch handlers
	engine.GET(ssh2.DefaultSSH2APIWebSocketURI, ssh2.NewWebsocketConnectionFunc)
	engine.GET(ssh2.DefaultSSH2APISessionQueryURI, ssh2.QuerySSH2SessionsFunc)
	engine.POST(ssh2.DefaultSSH2APISessionAddURI, ssh2.AddSSH2SessionFunc)
	engine.POST(ssh2.DefaultSSH2APISessionDeleteURI, ssh2.DeleteSSH2SessionFunc)
	engine.POST(ssh2.DefaultSSH2APISessionCloseURI, ssh2.CloseSSH2SessionFunc)

}
