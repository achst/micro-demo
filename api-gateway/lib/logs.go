package lib

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

// InitLogger init Logger
func InitLogger(path string, separates []string, level, maxdays int) *logs.BeeLogger {
	logger := logs.NewLogger()
	var confMap = map[string]interface{}{
		"filename": path,
		"maxdays":  maxdays,
		"daily":    true,
		"rotate":   true,
		"level":    level,
		"separate": separates,
	}
	conf, _ := json.Marshal(confMap)
	// 多文件输出和日志切分
	logger.SetLogger(logs.AdapterMultiFile, string(conf))
	// 控制台输出
	logger.SetLogger(logs.AdapterConsole)
	// 输出文件名和行号
	logger.EnableFuncCallDepth(true)
	// 异步输出, 缓冲 chan 的大小 1e3
	logger.Async(1e3)
	return logger
}
