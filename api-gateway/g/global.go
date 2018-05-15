package g

import (
	"github.com/hopehook/micro-demo/api-gateway/lib"

	"github.com/astaxie/beego/logs"
)

// Conf is to read demo.conf
var Conf *lib.Config

// Logger is global log fd
var Logger *logs.BeeLogger
