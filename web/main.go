package main

import (
	"log"
	"web/middle_ware"
	all_router "web/router"

	"github.com/gin-gonic/gin"

	"github.com/micro/go-web"
)

func main() {
	router := gin.Default()

	//使用全局中间件,跨域请求
	router.Use(middle_ware.CrosMiddleWare)
	//注册路由组
	all_router.InitRouter(router)

	service := web.NewService(
		web.Name("xlf.web.web.web"),
		web.Version("latest"),
		web.Handler(router),
		web.Address(":8081"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
