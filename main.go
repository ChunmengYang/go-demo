package fabric

import (
	"flag"
	"go-demo/g"
	"go-demo/http"
	"go-demo/socket"
	"go-demo/rpc"
)

func main()  {
	cfg := flag.String("c", "cfg.json", "configuration file")

	flag.Parse()

	g.ParseConfig(*cfg)
	g.InitLog()

	go http.Start()

	go socket.Start()

	go rpc.Start()

	select {}
}
