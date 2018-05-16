# micro-demo

A go-micro demo

### How to run it?

* run consul

    consul agent -dev

* run micro api

    micro api --handler=proxy

* run api-gateway with conf

    go run main.go api_gateway.conf

* run service-order

    go run main.go

* look in browser

    http://localhost:8080/api/order/list

* run micro web

    micro web

* look in browser

    http://localhost:8082