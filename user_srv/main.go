package main

import (
	"user_srv/controller"
	_ "user_srv/data_source"

	admin_user "user_srv/proto/admin_user"
	front_user "user_srv/proto/front_user"

	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("xlf.user.srv.user_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//example.RegisterExampleHandler(service.Server(), new(handler.Example))
	front_user.RegisterFrontUserHandler(service.Server(), new(controller.FrontUser))
	admin_user.RegisterAdminUserHandler(service.Server(), new(controller.AdminUser))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("xlf.user.srv.user_srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("xlf.user.srv.user_srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
