package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/FamousMai/go_study/grpc_mym/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// SayHi 挂载方法
func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(ctx)
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888") //这里不要写成rpc了
	if err != nil {
		fmt.Println("报错了")
	}

	// 创建服务
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})

	// 创建监听
	errS := s.Serve(l)
	if errS != nil {
		fmt.Println("失败：" + errS.Error())
	}
	fmt.Println("成功！")
}
