package rpc

import (
	"log"
	"testing"
	"github.com/toolkits/net"
	"time"
	"fmt"
)

func TestStart(t *testing.T) {
	client, err :=net.JsonRpcClient("tcp","127.0.0.1:1990", time.Second * 30)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &AgentRequest{sessionId: "testing"}
	var reply AgentResponse

	err = client.Call("Agent.GetDate", args, &reply)
	if err != nil {
		log.Fatal("call error:", err)
	}

	fmt.Printf("Current Time: %d", reply.date)
}