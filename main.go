package main

import (
	"github.com/lvboda/quick-chat/middleware"
	"github.com/lvboda/quick-chat/router"
	"github.com/lvboda/quick-chat/utils"
)

func main() {
	bootstrap()
}

func bootstrap() {
	app := utils.CreateApp()

	middleware.RegisterMiddleware(app)
	router.RegisterRoutes(app)

	app.RunTLS(utils.GetConfig().Server.Port, utils.CertFilePath, utils.KeyFilePath)
}
