package ports

import (
	"context"
	"github.com/AfRpEng/gorder-v2/common/genproto/orderpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct{}

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

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
