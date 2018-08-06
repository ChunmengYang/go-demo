package main

import (
	"flag"
	"github.com/ChunmengYang/go-demo/g"
	"github.com/ChunmengYang/go-demo/http"
		)

func main()  {
	cfg := flag.String("c", "cfg.json", "configuration file")

	flag.Parse()

	g.ParseConfig(*cfg)
	g.InitLog()

	go http.Start()

	//go socket.Start()

	//go rpc.Start()

	select {}
}
