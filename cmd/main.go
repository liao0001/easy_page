package main

import (
	"flag"
	"fmt"
	"github.com/liao0001/easy_page"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "conf.yml", "config path")
}

func main() {
	flag.Parse()
	conf, err := easy_page.LoadConfig(configPath)
	if err != nil {
		log.Fatalln("读取配置失败:", err.Error())
	}

	fmt.Println(conf)
}
