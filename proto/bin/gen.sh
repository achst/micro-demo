export GOPATH=/Users/tanshuai/lab/go

cd $GOPATH/src/github.com/hopehook/micro-demo/proto/service-order

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./service_order.proto