package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"jasontest/handler"
	"jasontest/subscriber"

	jasontest "jasontest/proto/jasontest"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.jasontest"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	jasontest.RegisterJasontestHandler(service.Server(), new(handler.Jasontest))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.jasontest", service.Server(), new(subscriber.Jasontest))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.jasontest", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
