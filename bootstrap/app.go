package bootstrap

import (
	"github.com/gin-gonic/gin"
	_ "github.com/zhouchang2017/epp/config"
	"github.com/zhouchang2017/epp/infrastructure"
	"github.com/zhouchang2017/epp/routes"

)

var App *app

type app struct {

	routers *gin.Engine
}

func newApp() *app {
	r := gin.Default()

	app := &app{
		routers: r,
	}

	infrastructure.Init()
	routes.ApiRouter(app.routers)

	app.routers.Run()

	return app
}

func Run() {
	App = newApp()
}
