package main

import (
	"github.com/AfRpEng/goorder-v2/order/app"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (s HTTPServer) GetCustomerCustomerIDOrders(c *gin.Context, customerID string) {

}

func (s HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {

}
