package socket

import (
	"net"
	"time"
	"fmt"
	"github.com/ChunmengYang/go-demo/g"
)

func Start() {
	if !g.Config().Socket.Enabled {
		return
	}

	ip := g.Config().Socket.IP
	port := g.Config().Socket.Port
	if port <= 0 {
		return
	}

	tcpAddr := net.TCPAddr{
		IP: net.ParseIP(ip),
		Port: port,
	}
	listener, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		g.Logger().Println(fmt.Sprintf("Fatal error: %s", err.Error()))
		return
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

var (
	clientConns = make(map[string]net.Conn)
)

func broadcast()  {
	msg := make([]byte, 128)
	for {
		//fmt.Print("Please input:")
		_, err := fmt.Scan(&msg)
		if err != nil {
			fmt.Println("Input error:", err.Error())
		}

		for _, conn := range clientConns  {
			conn.Write(msg)
		}
	}
}

func handleClient(conn net.Conn) {
	key := conn.RemoteAddr().String()
	clientConns[key] = conn

	defer func() {
		conn.Close()
		delete(clientConns, key)
	}()

	data := make([]byte, 128)
	for {
		i, err := conn.Read(data)
		if err != nil {
			break
		}
		daytime := time.Now().Format("15:04:05")
		fmt.Println(fmt.Sprintf("%s(%s):%s", key, daytime, string(data[0:i])))
	}

	fmt.Println(fmt.Sprintf("%s: Disconnect", key))
}