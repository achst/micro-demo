package g

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hopehook/micro-demo/api-gateway/lib"
	"github.com/hopehook/micro-demo/proto/service-order"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	//"github.com/micro/go-micro/transport"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
)

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// trace wrapper attaches a unique trace ID - timestamp
type traceWrapper struct {
	client.Client
}

// Implements client.Wrapper as traceWrapper
func traceWrap(c client.Client) client.Client {
	return &traceWrapper{c}
}

func (t *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = map[string]string{}
	}
	md["X-Trace-Id"] = fmt.Sprintf("%d", time.Now().Unix())
	ctx = metadata.NewContext(ctx, md)
	return t.Client.Call(ctx, req, rsp)
}

// TODO
func metricsWrap(cf client.CallFunc) client.CallFunc {
	return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := cf(ctx, addr, req, rsp, opts)
		fmt.Printf("[Metrics Wrapper] called: %s %s.%s duration: %v\n", addr, req.Service(), req.Method(), time.Since(t))
		return err
	}
}

// InitGlobal init glabal varibles
func InitGlobal() {
	// handler Logs
	{
		section := "HandlerLogs"
		path := Conf.Get(section, "path")
		separate := Conf.Get(section, "separate")
		separates := strings.Split(separate, ",")
		level, _ := strconv.Atoi(Conf.Get(section, "level"))
		maxdays, _ := strconv.Atoi(Conf.Get(section, "maxdays"))
		Logger = lib.InitLogger(path, separates, level, maxdays)
	}
}

// InitRPC init micro rpc clients
func InitRPC() {
	roundRobinWrap := roundrobin.NewClientWrapper()
	// set DefaultClient Option
	client.DefaultClient = client.NewClient(
		client.RequestTimeout(time.Second*5),
		client.DialTimeout(time.Second*5),
		client.Wrap(roundRobinWrap), // using a round robin client wrapper
		client.Wrap(traceWrap),
		client.Wrap(logWrap),
		client.WrapCall(metricsWrap),
		//client.Transport(
		//	transport.NewTransport(transport.Secure(true)),
		//),
	)
	// use the generated client stub
	OrderServiceClient = service_order.NewOrderServiceClient("go.micro.srv.order", client.DefaultClient)
}
