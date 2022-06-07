package rounter

import (
	"fmt"
	"github.com/eminoz/todo/controller"
	api "github.com/eminoz/todo/pb"
	"github.com/eminoz/todo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func BaseRpc() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()

	orderCollectionSetting := service.OrderCollectionSetting()
	todoCollectionSetting := service.TodoCollectionSetting()

	user := controller.User{UserServices: orderCollectionSetting}
	todo := controller.Todo{TodoService: todoCollectionSetting}

	api.RegisterCustomerServiceServer(newServer, &user)
	api.RegisterTodoServiceServer(newServer, &todo)

	reflection.Register(newServer)
	err = newServer.Serve(listen)
	if err != nil {
		panic(err)
	}
	fmt.Print("server stared on port 4040 ")
}
