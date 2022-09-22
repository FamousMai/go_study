package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/FamousMai/go_study/grpc_mym/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Println(err)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭失败" + err.Error())
		}
	}(conn)
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我从客户端来"})
	fmt.Println(req.GetMessage())
}
