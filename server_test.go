package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

var client *rpc.Client

func TestServer(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	client := jsonrpc.NewClient(conn)

	var (
		sum  int
		wsum string
		rsum int
	)

	elem := ArgsSum{
		Item1: 2,
		Item2: 3,
	}

	if err := client.Call("MyServer.Sum", elem, &sum); err != nil {
		t.Error(err)
	} else if sum != elem.Item1+elem.Item2 {
		t.Errorf("Expected : <%+v>,received: <%+v>", elem.Item1+elem.Item2, sum)
	}

	w := ArgsWrite{sum, "/home/silviu/go/jsonrpc1/h.txt"}

	if err := client.Call("MyServer.Write", w, &wsum); err != nil {
		t.Error(err)
	} else if wsum != "sum has been written to file" {
		t.Errorf("Expected : <%+v>,received: <%+v>", "sum has been written to file", wsum)
	}

	r := ArgsRead{w.FilePath, sum}

	if err := client.Call("MyServer.Read", r, &rsum); err != nil {
		t.Error(err)
	} else if rsum != elem.Item1+elem.Item2 {
		t.Errorf("Expected : <%+v>,received: <%+v>", elem.Item1+elem.Item2, rsum)
	}

}
