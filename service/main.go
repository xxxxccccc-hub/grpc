package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloRPCServer
}

func (s *server ) SayHi(ctx context.Context, req *hello_grpc.Req) (*hello_grpc.Res, error)   {
	fmt.Printf(req.GetMessage());
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc内容"},nil

}

func main()  {
	l,_ :=net.Listen("tcp",":8888")
	s:=grpc.NewServer()
	hello_grpc.RegisterHelloRPCServer(s,&server{})
	s.Serve(l)
}
