package main

import (
	"log"
	"net"

	"github.com/moricho/whois/service/greeter"
	pb "github.com/moricho/whois/service/proto"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	greeterService := &greeter.GreeterService{}

	pb.RegisterGreeterServer(server, greeterService)
	server.Serve(listenPort)
}
