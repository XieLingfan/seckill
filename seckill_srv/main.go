package main

import (
	"seckill_srv/controller"

	"github.com/micro/go-grpc"

	seckill "seckill_srv/proto/seckill"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("xlf.seckill.srv.seckill_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	seckill.RegisterSecKillHandler(service.Server(), new(controller.SecKill))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
