package main

import (
	"fmt"
	"path/filepath"

	"github.com/lvboda/quick-chat/middleware"
	"github.com/lvboda/quick-chat/router"
	"github.com/lvboda/quick-chat/utils"
)

func bootstrap() {
	config := utils.GetConfig()
	app := utils.CreateApp()

	app.Static("/assets", filepath.Join(utils.GetExecDirPath(), "../assets"))
	middleware.RegisterMiddleware(app)
	router.RegisterRoutes(app)

	addr := fmt.Sprintf(":%s", config.Server.Port)
	// app.RunTLS(addr, path.Join(utils.GetExecDirPath(), "../config/tls.pem"), path.Join(utils.GetExecDirPath(), "../config/tls.key"))
	app.Run(addr)
}

func main() {
	bootstrap()
}
