package socket

import (
    "fmt"
    "net"
    "testing"
)

const (
    addr = "127.0.0.1:1989"
)

var (
    msg = make(chan string)
)

func TestStart(t *testing.T) {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("Failed to connect to the server:", err.Error())
        return
    }
    fmt.Println("Connected")

    go showMessage(conn)

    sendMessage(conn)
}


func sendMessage(conn net.Conn) {
    defer conn.Close()
    for {
        temp := <- msg
        conn.Write([]byte(fmt.Sprintf("%s answer", temp)))
    }
}

func showMessage(conn net.Conn) {
    buf := make([]byte, 128)
    for {
        c, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Read server data error:", err.Error())
        }
        temp := string(buf[0:c])

        msg <- temp
        fmt.Println(temp)
    }
}