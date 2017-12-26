package test

import (
	"net"
	"net/rpc/jsonrpc"
	m "github.com/AAGerasimovich/test_rpc_go/model"

	"testing"
)

func TestSet(t *testing.T) {
	var reply m.Reply
	args := m.Args{
		Name: "Log",
	}
	reply, err := client("Set", args)
	if reply.Name != args.Name {
		t.Fatalf("%v", err)
	}
}

func TestGet(t *testing.T) {
	var reply m.Reply
	args := m.Args{
		Name: "Log",
	}
	reply, _ = client("Set", args)
	args.UUID = reply.UUID
	reply, err := client("Get", args)
	if reply.UUID != args.UUID {
		t.Fatalf("%v", err)
	}
}

func client(method string, args m.Args) (m.Reply, error) {
	conn, err := net.Dial("tcp", "localhost:8222")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := jsonrpc.NewClient(conn)

	var reply m.Reply

	err = c.Call("Model."+method, args, &reply)

	if err != nil {
		return reply, err
	}
	return reply, nil
}
