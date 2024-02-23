package main

import (
	"fmt"
	"log"
	"net"

	"github.com/allan7yin/grpc-go/internal/config"

	"google.golang.org/grpc"
)

func main() {
	c := config.ViperConfigInit()
	srvConf := c.GetServerConfig()
	listener, err := net.Listen(srvConf.GrpcProtocol, fmt.Sprintf("%s:%s", srvConf.Host, srvConf.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	fmt.Println("Server is running on port:", srvConf.GrpcPort)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
