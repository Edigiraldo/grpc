package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/Edigiraldo/grpc/exampb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer clientConn.Close()

	examServiceClient := exampb.NewExamServiceClient(clientConn)

	exam := exampb.Exam{Id: "e1", Name: "exam1"}
	SetExam(examServiceClient, &exam)
	GetExam(examServiceClient, exam.Id)

	questions := []*exampb.Question{
		{
			Id:       "q1e1",
			ExamId:   "e1",
			Question: "'OS' computer abbreviation usually means?",
			Answer:   "Operating System",
		},
		{
			Id:       "q2e1",
			ExamId:   "e1",
			Question: "How many Gigabytes are in 1 Terabyte??",
			Answer:   "1024",
		},
		{
			Id:       "q3e1",
			ExamId:   "e1",
			Question: "What is the name of the address of each computer on the Internet?",
			Answer:   "IP Address",
		},
		{
			Id:       "q4e1",
			ExamId:   "e1",
			Question: "What command can you use to count the number of lines, words, and characters in a file?",
			Answer:   "wc",
		},
	}
	SetQuestions(examServiceClient, questions)

	studentIds := []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9"}
	EnrollStudents(examServiceClient, exam.Id, studentIds)
	GetStudentsPerExam(examServiceClient, exam.Id)

	TakeExam(examServiceClient, exam.Id)
}

func SetExam(examServiceClient exampb.ExamServiceClient, exam *exampb.Exam) {
	log.Println("<----------------------SetExam---------------------->")
	defer log.Println("<----------------------SetExam---------------------->")

	_, err := examServiceClient.SetExam(context.Background(), exam)
	if err != nil {
		log.Fatalf("error while setting exam %s: %v\n", exam.Id, err)
	}
	log.Printf("successfully set exam: %v\n", exam.Name)
}

func GetExam(examServiceClient exampb.ExamServiceClient, examId string) {
	log.Println("<----------------------GetExam---------------------->")
	defer log.Println("<----------------------GetExam---------------------->")
	req := exampb.GetExamRequest{
		Id: examId,
	}

	res, err := examServiceClient.GetExam(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while getting exam %s: %v\n", examId, err)
	}

	log.Printf("exam: %v\n", res)
}

func SetQuestions(examServiceClient exampb.ExamServiceClient, questions []*exampb.Question) {
	log.Println("<----------------------SetQuestions---------------------->")
	defer log.Println("<----------------------SetQuestions---------------------->")
	stream, err := examServiceClient.SetQuestions(context.Background())
	if err != nil {
		log.Fatalf("error while calling SetQuestions: %v\n", err)
	}

	for _, question := range questions {
		log.Printf("sending question: %s\n", question.Id)
		err = stream.Send(question)
		if err != nil {
			log.Fatalf("error while sending question %s: %v\n", question.Id, err)
		}
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v\n", err)
	}

	log.Printf("response from server: %v\n", msg)
}

func EnrollStudents(examServiceClient exampb.ExamServiceClient, examId string, studentIds []string) {
	log.Println("<----------------------EnrollStudents---------------------->")
	defer log.Println("<----------------------EnrollStudents---------------------->")

	stream, err := examServiceClient.EnrollStudents(context.Background())
	if err != nil {
		log.Fatalf("error while calling EnrollStudents: %v\n", err)
	}

	for _, studentId := range studentIds {
		req := exampb.EnrollStudentsRequest{
			StudentId: studentId,
			ExamId:    examId,
		}
		log.Printf("enrolling student %s in exam %s\n", studentId, examId)
		err = stream.Send(&req)
		if err != nil {
			log.Fatalf("error while enrolling student %s: %v\n", studentId, err)
		}
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v\n", err)
	}

	log.Printf("response from server: %v\n", msg)
}

func GetStudentsPerExam(examServiceClient exampb.ExamServiceClient, examId string) {
	log.Println("<----------------------GetStudentsPerExam---------------------->")
	defer log.Println("<----------------------GetStudentsPerExam---------------------->")
	req := exampb.GetStudentsPerExamRequest{
		ExamId: examId,
	}
	stream, err := examServiceClient.GetStudentsPerExam(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while sending calling GetStudentsPerExam: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading from server: %v\n", err)
		}
		log.Printf("student %s with name %s and age %d is enrolled in exam %v\n", msg.GetId(), msg.GetName(), msg.GetAge(), examId)
	}
}

func TakeExam(examServiceClient exampb.ExamServiceClient, examId string) {
	log.Println("<----------------------TakeExam---------------------->")
	defer log.Println("<----------------------TakeExam---------------------->")
	req := exampb.TakeExamRequest{
		ExamId: examId,
		Answer: "2000",
	}
	stream, err := examServiceClient.TakeExam(context.Background())
	if err != nil {
		log.Fatalf("error while calling TakeExam: %v\n", err)
	}
	err = stream.Send(&req)
	if err != nil {
		log.Fatalf("error while sending request: %v\n", err)
	}

	numberOfQuestions := 4
	waitChannel := make(chan struct{})
	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving data from TakeExam: %v\n", err)
				break
			}
			log.Printf("response from TakeExam: %v\n", res)
		}
		close(waitChannel)
	}()
	<-waitChannel
}
