package main

import (
	"net"

	rpc "form-manager/pkg/server"
	pb "form-manager/pkg/server/pb"

	"google.golang.org/grpc"
)

func main() {

	tcp, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterFormServer(server, &rpc.FormMangement{})
	
	if err := server.Serve(tcp); err != nil {
		panic(err)
	}

}
