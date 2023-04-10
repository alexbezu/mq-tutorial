package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/alexbezu/mq-tutorial/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMQserviceServer
	mux sync.Mutex
	mq  map[string]chan []byte
}

func main() {
	s := &server{}
	listener, err := net.Listen("tcp", "localhost:3322")
	if err != nil {
		fmt.Errorf("Error: ", err)
		return
	}

	// var opts []grpc.ServerOption
	grpcServer := grpc.NewServer()
	pb.RegisterMQserviceServer(grpcServer, s)
	grpcServer.Serve(listener)
}

func (s *server) MQput(ctx context.Context, req *pb.MQputRequest) (*pb.MQputReply, error) {
	s.mux.Lock()
	if s.mq == nil {
		s.mq = make(map[string]chan []byte)
	}
	_, ok := s.mq[req.Qname]
	if !ok {
		s.mq[req.Qname] = make(chan []byte, 64)
	}
	s.mux.Unlock()
	s.mq[req.Qname] <- req.Value
	return &pb.MQputReply{Value: []byte("Ok")}, nil
}

func (s *server) MQpop(ctx context.Context, req *pb.MQpopRequest) (*pb.MQpopReply, error) {
	s.mux.Lock()
	if s.mq == nil {
		s.mq = make(map[string]chan []byte)
	}
	_, ok := s.mq[req.Qname]
	if !ok {
		s.mq[req.Qname] = make(chan []byte, 64)
	}
	s.mux.Unlock()
	reply := pb.MQpopReply{Value: <-s.mq[req.Qname]}
	return &reply, nil
}
