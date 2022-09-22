package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/FamousMai/go_study/grpc_mym/pb"
	person_grpc "github.com/FamousMai/go_study/grpc_mym/pb/person"
	"google.golang.org/grpc"
	"net"
	"time"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

type personServer struct {
	person_grpc.UnimplementedSearchServiceServer
}

// SayHi 挂载方法
func (s *server) SayHi(context.Context, *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

// Search 普通请求
func (s *personServer) Search(_ context.Context, req *person_grpc.PersonReq) (res *person_grpc.PersonRes, err error) {
	fmt.Println(req.Name)
	res = &person_grpc.PersonRes{Name: "我收到了" + req.Name + "的信息"}
	return res, nil
}

// SearchIn 流式请求
func (s *personServer) SearchIn(server person_grpc.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil { //有可能传完了，有可能传错了
			err := server.SendAndClose(&person_grpc.PersonRes{Name: "完成了" + err.Error()})
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}
	return nil
}

// SearchOut 流式返回
func (s *personServer) SearchOut(req *person_grpc.PersonReq, server person_grpc.SearchService_SearchOutServer) error {
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		server.Send(&person_grpc.PersonRes{Name: "我拿到了" + req.Name})
		i++
	}
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
