package main

import (
	"context"
	"github.com/AfRpEng/goorder-v2/common/config"
	"github.com/AfRpEng/goorder-v2/common/genproto/orderpb"
	"github.com/AfRpEng/goorder-v2/common/server"
	"github.com/AfRpEng/goorder-v2/order/ports"
	"github.com/AfRpEng/goorder-v2/order/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	go server.RunGRPCServer(serviceName, func(s *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		orderpb.RegisterOrderServiceServer(s, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{
			app: application,
		}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
