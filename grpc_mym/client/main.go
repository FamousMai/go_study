package main

import (
	"context"
	"fmt"
	person_grpc "github.com/FamousMai/go_study/grpc_mym/pb/person"
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

	client := person_grpc.NewSearchServiceClient(conn)
	res, _ := client.Search(context.Background(), &person_grpc.PersonReq{Name: "我是老麦"})
	fmt.Println(res.GetName())
}
