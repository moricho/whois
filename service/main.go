package main

import (
	"log"
	"net"

	pb "github.com/moricho/whois/service/greeter/proto"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	greeterService := &greeter.GreeterService{}

	pb.RegisterGreeterService(server, greeterService)
	server.Serve(listenPort)
}
