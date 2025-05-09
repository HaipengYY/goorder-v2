package main

import (
	"context"
	"github.com/AfRpEng/goorder-v2/common/config"
	"github.com/AfRpEng/goorder-v2/common/genproto/stockpb"
	"github.com/AfRpEng/goorder-v2/common/server"
	"github.com/AfRpEng/goorder-v2/stock/ports"
	"github.com/AfRpEng/goorder-v2/stock/service"
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
	serviceName := viper.GetString("stock.service_name")
	serverType := viper.GetString("stock.server-to-run")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		//暂时不用
	default:
		logrus.Fatal("stock server type not supported")
	}
}
