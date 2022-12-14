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
	SetQuestion(ctx context.Context, question *models.Question) error
	SetEnrollment(ctx context.Context, enroll *models.Enrollment) error
	GetStudentsPerExam(ctx context.Context, examId string) ([]models.Student, error)
	GetQuestionsPerExam(ctx context.Context, examId string) ([]models.Question, error)
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

func SetEnrollment(ctx context.Context, enroll *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enroll)
}

func GetStudentsPerExam(ctx context.Context, examId string) ([]models.Student, error) {
	return implementation.GetStudentsPerExam(ctx, examId)
}

func GetQuestionsPerExam(ctx context.Context, examId string) ([]models.Question, error) {
	return implementation.GetQuestionsPerExam(ctx, examId)
}
