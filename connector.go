package main

import (
	"context"
	"log"
	"time"

	// "github.com/codenotary/immudb/pkg/api/schema"
	schema "github.com/alexbezu/mq-tutorial/pb"
	"google.golang.org/grpc"
)

//TODO:
//go:generate PATH="$PATH:$(go env GOPATH)/bin" protoc --go_out=. --go-grpc_out=. pb/mq.proto

var conn *grpc.ClientConn
var err error
var c schema.MQserviceClient
var ctx context.Context
var cancel context.CancelFunc

func connect2db() {
	log.Print("Trying to connect to MQ: ")
	conn, err = grpc.Dial("localhost:3322", grpc.WithInsecure(), grpc.WithReturnConnectionError())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Print("Connected!")
	}

	c = schema.NewMQserviceClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
}

func closedb() {
	conn.Close()
	cancel()
}
