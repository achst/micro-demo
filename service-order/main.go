package main

import (
	"context"
	"log"
	"time"

	"github.com/hopehook/micro-demo/proto/service-order"
	"github.com/hopehook/micro-demo/service-order/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	//"github.com/micro/go-micro/transport"
)

const microName = "go.micro.srv.order"

// implements the server.HandlerWrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[%v] server request: %s\n", time.Now(), req.Method())
		return fn(ctx, req, rsp)
	}
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(microName),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		// wrap the handler
		micro.WrapHandler(logWrapper),
		// setup a new transport with secure option
		//micro.Transport(
		//	// create new transport
		//	transport.NewTransport(
		//		// set to automatically secure
		//		transport.Secure(true),
		//	),
		//),
	)

	// Init will parse the command line flags.
	// optionally setup command line usage
	service.Server().Init(
		server.Wait(true),
	)
	//service.Init()

	// Register handler
	service_order.RegisterOrderServiceHandler(service.Server(), new(handler.OrderHandler))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
