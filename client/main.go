package main

import (
	"context"
	"fmt"
	"github.com/xxxxccccc-hub/grpc/pb"
	"google.golang.org/grpc"
)

func main()  {
	conn,e := grpc.Dial("localhost:8888",grpc.WithInsecure())
	fmt.Println(e)
	defer conn.Close()
	client := hello_grpc.NewHelloRPCClient(conn)
	req,_ := client.SayHi(context.Background(),&hello_grpc.Req{Message:"我从客户端来"})
	fmt.Println(req.GetMessage())

}
