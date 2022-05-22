package main

import (
	"github.com/tyokoyama/golangcafe/evans/protosample"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	protosample.SayHello()

	reflection.Register(s)
}