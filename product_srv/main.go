package main

import (
	"product_srv/controller"
	_ "product_srv/data_source"

	product "product_srv/proto/product"
	seckill "product_srv/proto/seckill"

	"github.com/micro/go-grpc"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("xlf.product.srv.product_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//example.RegisterExampleHandler(service.Server(), new(handler.Example))

	product.RegisterProductsHandler(service.Server(), new(controller.Products))
	seckill.RegisterSecKillsHandler(service.Server(), new(controller.SecKills))
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("xlf.product.srv.product_srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("xlf.product.srv.product_srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
