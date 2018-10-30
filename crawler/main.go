package main

import (
	// net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
	"learn-go/crawler/engine"
	"learn-go/crawler/zhenai/parser"
)

func main () {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}