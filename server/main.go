package main

import (
	"context"
	"github.com/rodbalbino/grpc_tutorial/proto"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main(){
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\033[1;36m%s\033[0m", "\nServer UP http://localhost:4040")

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context,request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
