package main

import (
	pb "grpc/proto"
)

type Person struct {
	ID          int32
	Name        string
	Email       string
	PhoneNumber string
}

var nextID int32 = 1
var persons = make(map[int32]Person)

type Server struct {
	pb.UnimplementedPersonServiceServer
}

func main() {

}
