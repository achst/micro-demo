package lib

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const delimiter = "$"

// Config struct
type Config struct {
	Dict    map[string]string
	section string
}

// InitConfig init config struct
func InitConfig(path string) (c *Config) {
	c = &Config{
		Dict:    make(map[string]string),
		section: "",
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// 忽略注释
		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}
		// 读取 section
		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.section = strings.TrimSpace(s[n1+1 : n2])
			continue
		}
		if len(c.section) == 0 {
			continue
		}
		// 读取 dict
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}
		key := c.section + delimiter + frist
		c.Dict[key] = strings.TrimSpace(second)
	}
	return
}

// Get config value
func (c *Config) Get(section, key string) (value string) {
	key = section + delimiter + key
	v, ok := c.Dict[key]
	if !ok {
		log.Panicln("Get config value failed, key: ", key)
	}
	return v
}
