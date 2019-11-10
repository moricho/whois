package hello

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/moricho/whois/service/hello/proto"
)

type HelloController struct{}

func (hc *HelloController) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if req.Name == "" {
		return nil, errors.New("error: empty input.")
	}
	msg := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloReply{
		Message: msg,
	}, nil
}
