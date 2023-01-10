package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/graiendor/anomaly_detector/internal"
	"github.com/graiendor/anomaly_detector/internal/anomalies"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"

	pb "github.com/graiendor/anomaly_detector/services"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3333, "The server port")
	k    = flag.Float64("k", 3.0, "STD anomaly coefficient")
)

func main() {
	flag.Parse()

	if *k < 0 {
		log.Fatalf("STD anomaly coefficient should be >= 0 (k = %f)", *k)
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := pb.NewTransmissionServiceClient(conn)
	in := &pb.Request{Id: 1}
	stream, err := client.FetchTransmission(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	var data []internal.Report
	for i := uint64(1); i != 0; i++ {
		resp, err := stream.Recv()
		if err == io.EOF {
			//done <- true
			return
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		data = append(data, internal.Report{
			SessionID: resp.SessionId,
			Frequency: resp.Frequency,
			Timestamp: resp.Timestamp,
		})

		if i%100 == 0 {
			anomalies.LogPredictions(data)
		}
		if i == 1000 {
			break
		}
	}
	anomalies.LogAnomalies(*k, data)
}
