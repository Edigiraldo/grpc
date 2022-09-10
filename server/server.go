package server

import (
	"context"

	"github.com/Edigiraldo/grpc/models"
	"github.com/Edigiraldo/grpc/repository"
	"github.com/Edigiraldo/grpc/studentpb"
)

type Server struct {
	repository repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repository: repo}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repository.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}
	err := s.repository.SetStudent(ctx, &student)
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}
