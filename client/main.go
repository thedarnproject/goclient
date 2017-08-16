package main

import (
	"flag"
	"log"

	"context"

	"github.com/Sirupsen/logrus"
	api "github.com/thedarnproject/thedarnapi/api"
	"google.golang.org/grpc"
)

func main() {
	apiAddress := flag.String("api", "localhost:8081", "address of gRPC API")
	logrus.Infof("connecting to server: %v", *apiAddress)
	clientConnection, err := grpc.Dial(*apiAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %v: %v", *apiAddress, err)
	}
	defer clientConnection.Close()

	client := api.NewErrorInClient(clientConnection)

	clientData := &api.Data{
		Error:    "this is the error",
		Platform: "linux",
		Plugin:   "go client",
	}

	logrus.Infof("sending data to server:\n%#v", *clientData)
	success, err := client.Submit(context.Background(), clientData)
	if err != nil {
		log.Fatalf("could not submit data to api server %v: %v", *clientData, err)
	}
	if success.Success {
		logrus.Info("server returned success")
	}
}
