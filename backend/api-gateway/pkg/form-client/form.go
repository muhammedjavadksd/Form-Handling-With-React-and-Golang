package formclient

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/form-client/form/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClient(cfg *config.Config) pb.FormClient {
	conn, err := grpc.NewClient(":"+cfg.FormGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		log.Fatalf(err.Error())
	}

	// defer conn.Close()
	return pb.NewFormClient(conn)

}
