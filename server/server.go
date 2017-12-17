package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/rpctests/server/proto"
)

type Server struct {
	count int
}

func (s *Server) Ping(ctx context.Context, in *pb.ServerRequest) (*pb.ServerResponse, error) {
	for i := 0; i < int(in.GetRqindex()); i++ {
		t, ok := ctx.Deadline()
		log.Printf("Counting %v with %v, %v", i, t, ok)
		time.Sleep(time.Millisecond * 100)
	}

	return &pb.ServerResponse{Rsindex: in.Rqindex + 1}, nil
}

func (s *Server) Serve(port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	server := grpc.NewServer(
		grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()),
	)

	pb.RegisterServerServer(server, s)
	log.Printf("Serving")
	server.Serve(lis)

	return nil
}

func main() {
	s := &Server{}
	s.Serve(50055)
}
