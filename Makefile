.PHONY: clean

MODULE=github.com/graiendor/anomaly_detector

build: build-server build-client

build-server:
	go build github.com/graiendor/anomaly_detector/cmd/server

build-client:
	go build github.com/graiendor/anomaly_detector/cmd/client

build-postgresql:
	go build github.com/graiendor/anomaly_detector/cmd/postgresql

clean:
	rm -rf client server