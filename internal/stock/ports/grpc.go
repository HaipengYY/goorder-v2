package ports

import (
	"context"
	"github.com/AfRpEng/goorder-v2/common/genproto/stockpb"
	"github.com/AfRpEng/goorder-v2/stock/app"
)

// app注入
type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (g GrpcServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
