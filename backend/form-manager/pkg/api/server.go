package api

import (
	"form-manager/pkg/config"
	rpc "form-manager/pkg/server"
	pb "form-manager/pkg/server/pb"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	GrpcPort string
	Engine   *grpc.Server
}

func InitGRPCServer(cfg *config.Config) *Server {

	server := grpc.NewServer()

	pb.RegisterFormServer(server, &rpc.FormMangement{})
	return &Server{
		GrpcPort: cfg.GrpcPort,
		Engine:   server,
	}

}

func (s *Server) Start() {
	tcp, err := net.Listen("tcp", ":"+s.GrpcPort)
	if err != nil {
		panic(err)
	}

	if err = s.Engine.Serve(tcp); err != nil {
		panic(err)
	}

}
