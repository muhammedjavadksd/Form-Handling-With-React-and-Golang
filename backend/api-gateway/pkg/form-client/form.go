package formclient

import (
	"api-gateway/pkg/form-client/form/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClient() pb.FormClient {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		log.Fatalf(err.Error())
	}

	// defer conn.Close()
	return pb.NewFormClient(conn)

}
