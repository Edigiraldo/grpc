package repository

import (
	"context"

	"github.com/Edigiraldo/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetExam(ctx context.Context, id string) (*models.Exam, error)
	SetExam(ctx context.Context, exam *models.Exam) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetExam(ctx context.Context, id string) (*models.Exam, error) {
	return implementation.GetExam(ctx, id)
}

func SetExam(ctx context.Context, exam *models.Exam) error {
	return implementation.SetExam(ctx, exam)
}
