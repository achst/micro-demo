package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/hopehook/micro-demo/api-gateway/g"
	"github.com/hopehook/micro-demo/api-gateway/lib"
	"github.com/hopehook/micro-demo/api-gateway/router"
	"github.com/hopehook/micro-demo/api-gateway/util"
	"github.com/micro/go-web"
)

const defaultConf = "api_gateway.conf"
const microName = "go.micro.api.api"

func init() {
	// 初始化配置文件
	g.Conf = lib.InitConfig(util.GetConfPath(defaultConf))
	// 初始化全局变量
	g.InitGlobal()
	// 初始化 rpc 客户端
	g.InitRPC()
	// 初始化基础性日志
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 初始化路由注册
	router.Init()
}

func main() {
	// http server
	srv := &http.Server{
		ReadTimeout:    15 * time.Second, // 读取 http 请求超时时间(header 和 body)
		WriteTimeout:   15 * time.Second, // 响应 http 返回超时时间
		MaxHeaderBytes: 1 << 20,          // 1 MB. 限制头部最大字节数
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefulShutdown := func() error {
		srv.Shutdown(ctx)
		return errors.New("server gracefully stopped")
	}

	// create service
	service := web.NewService(
		web.Name(microName),
		web.Version("v1"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
		web.Server(srv),                  // set custom server
		web.BeforeStop(gracefulShutdown), // graceful shutdown
	)
	service.Init()                     // init service
	service.Handle("/", router.Router) // Register Handler
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
