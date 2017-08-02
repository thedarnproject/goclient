package main

import (
	api "github.com/thedarnproject/thedarnapi/api"

	"flag"
	"log"

	"google.golang.org/grpc"
	"context"
	"fmt"
)

func main() {
	apiAddress := flag.String("api", "localhost:8080", "address of gRPC API")
	clientConnection, err := grpc.Dial(*apiAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %v: %v", *apiAddress, err)
	}
	defer clientConnection.Close()

	client := api.NewErrorInClient(clientConnection)

	clientData := &api.Data{
		Error: "this is the error",
		Platform: "linux",
		Plugin: "go client",
	}

	success, err := client.Submit(context.Background(), clientData)
	if err != nil {
		log.Fatalf("could not submit data to api server %v: %v", *clientData, err)
	}
	fmt.Println(success.Success)

}
