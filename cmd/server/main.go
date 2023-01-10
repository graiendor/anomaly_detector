package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/graiendor/anomaly_detector/internal/sampler"
	pb "github.com/graiendor/anomaly_detector/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	port = flag.Int("port", 3333, "The server port")
)

type Transmission struct {
	pb.UnimplementedTransmissionServiceServer
}

func (s *Transmission) FetchTransmission(in *pb.Request, srv pb.TransmissionService_FetchTransmissionServer) error {
	sessionID := uuid.NewString()
	stats := sampler.GenerateStats()
	for {
		resp := pb.Response{
			SessionId: sessionID,
			Frequency: stats.GetSample(),
			Timestamp: time.Now().UTC().String(),
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
			return nil
		}

		log.Printf("Session id transmitted:\t%s", resp.SessionId)
		log.Printf("Freq transmitted:\t%f", resp.Frequency)
		log.Printf("Timestamp transmitted:\t%s\n\n", resp.Timestamp)
		//time.Sleep(time.Second / 10)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransmissionServiceServer(s, &Transmission{})
	log.Printf("start server on port: %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
