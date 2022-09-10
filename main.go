package main

import (
	"log"
	"net"

	"github.com/Edigiraldo/grpc/database"
	"github.com/Edigiraldo/grpc/server"
	"github.com/Edigiraldo/grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatalf("Error creating repository: %s", err.Error())
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
