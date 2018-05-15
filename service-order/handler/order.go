package handler

import (
	"context"
	"time"

	"log"

	"github.com/hopehook/micro-demo/proto/service-order"
	"github.com/hopehook/micro-demo/service-order/model"
	"github.com/micro/go-micro/metadata"
)

type OrderHandler struct{}

func (*OrderHandler) GetOrders(ctx context.Context, req *service_order.GetOrdersRequest, rsp *service_order.GetOrdersResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		rsp.Orders = nil
		return nil
	}
	orders, err := model.GetOrders(req.PageIndex, req.PageCount)
	if err != nil {
		return err
	}
	log.Println(md)
	rsp.Orders = orders
	time.Sleep(1 * time.Second)
	return nil
}
