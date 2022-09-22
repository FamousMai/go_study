package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/FamousMai/go_study/grpc_mym/pb"
	person_grpc "github.com/FamousMai/go_study/grpc_mym/pb/person"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

type personServer struct {
	person_grpc.UnimplementedSearchServiceServer
}

// SayHi 挂载方法
func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(ctx)
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

func (s *personServer) Search(ctx context.Context, req *person_grpc.PersonReq) (res *person_grpc.PersonRes, err error) {
	fmt.Println(ctx)
	name := req.GetName()
	res = &person_grpc.PersonRes{Name: "我收到了" + name + "的信息"}
	return res, nil
}
func (s *personServer) SearchIn(person_grpc.SearchService_SearchInServer) error {
	return nil
}
func (s *personServer) SearchOut(*person_grpc.PersonReq, person_grpc.SearchService_SearchOutServer) error {
	return nil
}
func (s *personServer) SearchIO(person_grpc.SearchService_SearchIOServer) error {
	return nil
}

func main() {
	fmt.Println(server{})

	l, err := net.Listen("tcp", ":8888") //这里不要写成rpc了
	if err != nil {
		fmt.Println("报错了")
	}

	// 创建服务
	s := grpc.NewServer()
	person_grpc.RegisterSearchServiceServer(s, &personServer{})

	// 创建监听
	errS := s.Serve(l)
	if errS != nil {
		fmt.Println("失败：" + errS.Error())
	}
	fmt.Println("成功！")
}
