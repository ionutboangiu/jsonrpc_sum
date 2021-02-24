package main

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type MyServer struct{}

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

func (srv *MyServer) Sum(args ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
	return nil
}

func (srv *MyServer) Write(args ArgsWrite, reply *string) error {
	file, err := os.Create(args.FilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer file.Close()

	b, err := json.Marshal(args)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	*reply = "sum has been written to file"

	return nil
}

func (srv *MyServer) Read(args ArgsRead, reply *int) error {
	file, err := os.Open(args.FilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	b := make([]byte, 150)

	nr, err := file.Read(b)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var m ArgsRead
	err = json.Unmarshal(b[:nr], &m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	*reply = args.Item

	return nil
}

func main() {
	var srv = new(MyServer)
	server := rpc.NewServer()

	err := server.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
