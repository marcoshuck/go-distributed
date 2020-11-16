package main

import (
	"fmt"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/cloud"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Initializing cloud provider microservice using gRPC")
	port := os.Getenv("CLOUD_PROVIDER_PORT")
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("Failed to initialize server on port:", port)
		panic(err)
	}

	s := grpc.NewServer()
	c := cloud.NewCloudServer()

	log.Println("Adding cloud provider microservice to register.")
	cloud.RegisterCloudServer(s, c)

	log.Println("Serving cloud provider microservice on port:", port)
	err = s.Serve(listener)
	if err != nil {
		log.Println("Failed to serve gRPC service on port:", port)
		panic(err)
	}
}
