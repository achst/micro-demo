package router

import (
	"github.com/hopehook/micro-demo/api-gateway/handler"
	"github.com/julienschmidt/httprouter"
)

var Router = httprouter.New()

// Init init routers
// url prefix must keep consistent with MicroName's last name.
// eg: url: "/api/*"  <->  "go.micro.api.api"
//     url: "/auth/*" <->  "go.micro.api.auth"
func Init() {
	Router.GET("/api/order/list", handler.Raw(handler.GetOrderList))
}
