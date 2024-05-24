package rpc

import (
	"context"
	"fmt"
	pb "form-manager/pkg/server/pb"
)

type FormMangement struct {
	pb.UnimplementedFormServer
}

func (f *FormMangement) InsertForm(ctx context.Context, in *pb.FormReq) (*pb.FormResponse, error) {
	fmt.Println("working")
	fmt.Println(in)
	return &pb.FormResponse{
		Id: 4,
	}, nil
}
