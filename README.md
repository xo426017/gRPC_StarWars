# Go gRPC Demo

## Prerequisites

* make
* protoc
* go 1.12+

## Commands

gen:
	protoc -I . --go_out="plugins=grpc:." starwars.proto

build:
	go build

run:
	go run starwars
