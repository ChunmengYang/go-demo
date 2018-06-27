package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Agent int

func (t *Agent) GetDate(args *AgentRequest, reply *AgentResponse) error {
	log.Println(args.sessionId)

	reply.date = time.Now().Unix()

	return nil
}

func Start() {
	addr := "127.0.0.1:1990"

	server := rpc.NewServer()
	server.Register(new(Agent))

	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("listener accept fail:", err)
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}