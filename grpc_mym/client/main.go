package main

import (
	"context"
	"fmt"
	person_grpc "github.com/FamousMai/go_study/grpc_mym/pb/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
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

	/**
	普通请求
	*/
	client := person_grpc.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &person_grpc.PersonReq{Name: "我是老麦"})
	if err != nil {
		fmt.Println("请求出错：" + err.Error())
	}
	fmt.Println(res.GetName())

	/**
	流式请求
	*/
	c, err := client.SearchIn(context.Background())
	if err != nil {
		fmt.Println("请求出错：" + err.Error())
	}
	i := 0
	for {
		if i > 10 {
			res, err := c.CloseAndRecv()
			fmt.Println(res, err)
			break
		}
		time.Sleep(1 * time.Second)
		err := c.Send(&person_grpc.PersonReq{Name: "我是 SearchIn 进来的信息"})
		if err != nil {
			fmt.Println(err)
		}
		i++
	}

}
