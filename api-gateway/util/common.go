package util

import (
	"os"
	"path/filepath"
)

// IsExist 文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// GetExecPath 获取执行路径
func GetExecPath() (path string, err error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetConfPath 获取配置文件路径
func GetConfPath(defaultConf string) (path string) {
	if len(os.Args) != 2 {
		// 默认配置文件路径
		// 当前可执行文件路径
		if excPath, err := GetExecPath(); err == nil {
			path = filepath.Join(excPath, defaultConf)
			if !IsExist(path) {
				// 当前GOPATH目录
				if cwd, err := os.Getwd(); err == nil {
					path = filepath.Join(cwd, defaultConf)
				}
			}
		}
		if !IsExist(path) {
			panic("Can not get default config path")
		}
	} else {
		// 用户指定配置文件路径
		path = os.Args[1]
		if !IsExist(path) {
			panic("Given config file path not exist")
		}
	}
	return
}
