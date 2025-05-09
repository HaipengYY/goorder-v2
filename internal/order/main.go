package main

import (
	"github.com/AfRpEng/gorder-v2/common/config"
	"github.com/AfRpEng/gorder-v2/common/genproto/orderpb"
	"github.com/AfRpEng/gorder-v2/common/server"
	"github.com/AfRpEng/gorder-v2/order/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	go server.RunGRPCServer(serviceName, func(s *grpc.Server) {
		svc := ports.NewGrpcServer()
		orderpb.RegisterOrderServiceServer(s, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
