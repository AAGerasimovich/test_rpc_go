package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	t "github.com/AAGerasimovich/test_rpc_go/model"
)

func main() {

	model := new(t.Model)

	server := rpc.NewServer()
	server.Register(model)

	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("Serving RPC handler")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		log.Println("Serving RPC handler")
		defer conn.Close()

	}

}
