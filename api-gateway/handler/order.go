package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/hopehook/micro-demo/api-gateway/g"
	service_order "github.com/hopehook/micro-demo/proto/service-order"
	"github.com/micro/go-micro/metadata"
)

// GetOrderList handler
func GetOrderList(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second*10)
	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := g.OrderServiceClient.GetOrders(ctx, &service_order.GetOrdersRequest{
		PageIndex: 0,
		PageCount: 10,
	})
	if err != nil {
		CommonWrite(w, r, -1, err.Error(), nil)
		return
	}
	CommonWriteSuccess(w, r, rsp.Orders)
}
