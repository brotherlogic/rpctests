package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/rpctests/server/proto"
)

//Call makes the call
func Call() {
	conn, err := grpc.Dial(":50055", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Printf("ERR %v", err)
	}
	client := pb.NewServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	log.Printf("CANCEL = %v", cancel)
	defer cancel()

	resp, err := client.Ping(ctx, &pb.ServerRequest{Rqindex: 100})

	log.Printf("%v and %v", resp, err)
	/*if err == nil {
	  monitor := s.monitorBuilder.NewMonitorServiceClient(conn)
	  messageLog := &pbd.MessageLog{Message: message, Entry: s.Registry}

	  defer cancel()
	  monitor.WriteMessageLog(ctx, messageLog, grpc.FailFast(false))
	  s.close(conn)
	}*/
}

func main() {
	Call()
}
