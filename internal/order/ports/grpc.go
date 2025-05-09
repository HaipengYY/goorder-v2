package ports

import (
	"context"
	"github.com/AfRpEng/goorder-v2/common/genproto/orderpb"
	"github.com/AfRpEng/goorder-v2/order/app"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	Application app.Application
}

func NewGrpcServer(application app.Application) *GrpcServer {
	return &GrpcServer{Application: application}
}

func (g GrpcServer) CreatOrder(ctx context.Context, request *orderpb.CreatOrderRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
