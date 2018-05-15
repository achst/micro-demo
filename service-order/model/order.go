package model

import (
	"github.com/hopehook/micro-demo/proto/service-order"
)

func GetOrders(pageIndex int32, pageCount int32) ([]*service_order.Order, error) {
	_, _ = pageIndex, pageCount
	orders := []*service_order.Order{
		{
			Id:         1,
			Code:       "",
			UserId:     1,
			Title:      "",
			Price:      9,
			CreateTime: "",
			UpdateTime: "",
		},
	}
	return orders, nil
}
