package servers

import (
	"context"
	"io"
	"log"

	"github.com/Edigiraldo/grpc/exampb"
	"github.com/Edigiraldo/grpc/models"
	"github.com/Edigiraldo/grpc/repository"
	"github.com/Edigiraldo/grpc/studentpb"
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

func (e *Exam) SetQuestions(stream exampb.ExamService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionsResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		question := models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			ExamId:   msg.GetExamId(),
		}

		err = e.repository.SetQuestion(stream.Context(), &question)
		if err != nil {
			return stream.SendAndClose(&exampb.SetQuestionsResponse{
				Ok: false,
			})
		}
	}
}

func (e *Exam) EnrollStudents(stream exampb.ExamService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.EnrollStudentsResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		enroll := models.Enrollment{
			StudentId: msg.GetStudentId(),
			ExamId:    msg.GetExamId(),
		}

		err = e.repository.SetEnrollment(stream.Context(), &enroll)
		if err != nil {
			return stream.SendAndClose(&exampb.EnrollStudentsResponse{
				Ok: false,
			})
		}
	}
}

func (e *Exam) GetStudentsPerExam(req *exampb.GetStudentsPerExamRequest, stream exampb.ExamService_GetStudentsPerExamServer) error {
	examId := req.GetExamId()
	students, err := e.repository.GetStudentsPerExam(context.TODO(), examId)
	if err != nil {
		return err
	}

	for _, s := range students {
		student := studentpb.Student{
			Id:   s.Id,
			Name: s.Name,
			Age:  s.Age,
		}
		if err = stream.Send(&student); err != nil {
			return err
		}
	}

	return nil
}

func (e *Exam) TakeExam(stream exampb.ExamService_TakeExamServer) error {
	answer, err := stream.Recv()
	if err == io.EOF {
		return nil
	}

	questions, err := e.repository.GetQuestionsPerExam(context.Background(), answer.GetExamId())
	if err != nil {
		return err
	}

	i := 0
	var currentQuestion = models.Question{}
	for {
		if i < len(questions) {
			currentQuestion = questions[i]
		}

		if i <= len(questions)-1 {
			questionToSend := &exampb.Question{
				Id:       currentQuestion.Id,
				Question: currentQuestion.Question,
			}
			err := stream.Send(questionToSend)
			if err != nil {
				log.Printf("Error sending question: %v", err)
				return err
			}
			i++
		}
		answer, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving answer: %v", err)
			return err
		}
		log.Println("Answer for question:", currentQuestion.Question, "is", answer.GetAnswer())
	}
}
