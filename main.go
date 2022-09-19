package main

import (
	"fmt"

	"github.com/lvboda/quick-chat/middleware"
	"github.com/lvboda/quick-chat/router"
	"github.com/lvboda/quick-chat/utils"
)

func bootstrap() {
	config := utils.GetConfig()
	app := utils.CreateApp()

	app.Static("/assets", "./assets")
	middleware.RegisterMiddleware(app)
	router.RegisterRoutes(app)

	addr := fmt.Sprintf(":%s", config.Server.Port)
	app.Run(addr)
}

func main() {
	bootstrap()
}
