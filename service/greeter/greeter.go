package greeter

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/moricho/whois/service/proto"
)

type GreeterService struct{}

func (gs *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if req.Name == "" {
		return nil, errors.New("error: empty input.")
	}
	msg := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloReply{
		Message: msg,
	}, nil
}
