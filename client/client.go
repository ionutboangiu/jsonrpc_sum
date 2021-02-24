package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type ArgsSum struct {
	Item1, Item2 int
}

type ArgsWrite struct {
	Item     int
	FilePath string
}

type ArgsRead struct {
	FilePath string
	Item     int
}

func main() {
	var (
		sum  int
		wsum string
		rsum int
	)

	client, _ := jsonrpc.Dial("tcp", "localhost:1234")

	elem := ArgsSum{2, 3}

	client.Call("MyServer.Sum", elem, &sum)
	fmt.Println(sum)

	w := ArgsWrite{sum, "/home/silviu/go/jsonrpc3/sum.json"}

	client.Call("MyServer.Write", w, &wsum)
	fmt.Println(wsum)

	r := ArgsRead{w.FilePath, sum}

	client.Call("MyServer.Read", r, &rsum)
	fmt.Printf("sum of %v and %v is %v\n", elem.Item1, elem.Item2, rsum)

}
