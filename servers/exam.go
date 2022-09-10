package servers

import (
	"context"

	"github.com/Edigiraldo/grpc/exampb"
	"github.com/Edigiraldo/grpc/models"
	"github.com/Edigiraldo/grpc/repository"
)

type Exam struct {
	repository repository.Repository
	exampb.UnimplementedExamServiceServer
}

func NewExamServer(repo repository.Repository) *Exam {
	return &Exam{repository: repo}
}

func (e *Exam) GetExam(ctx context.Context, req *exampb.GetExamRequest) (*exampb.Exam, error) {
	exam, err := e.repository.GetExam(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &exampb.Exam{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil
}

func (e *Exam) SetExam(ctx context.Context, req *exampb.Exam) (*exampb.SetExamResponse, error) {
	exam := models.Exam{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	err := e.repository.SetExam(ctx, &exam)
	if err != nil {
		return nil, err
	}

	return &exampb.SetExamResponse{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil
}
