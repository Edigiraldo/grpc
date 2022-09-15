package main

import (
	"context"
	"log"

	"github.com/Edigiraldo/grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial("localhost:5060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	studentServiceClient := studentpb.NewStudentServiceClient(clientConn)

	students := []*studentpb.Student{
		{Id: "s1", Name: "student1", Age: 1},
		{Id: "s2", Name: "student2", Age: 2},
		{Id: "s3", Name: "student3", Age: 3},
		{Id: "s4", Name: "student4", Age: 4},
		{Id: "s5", Name: "student5", Age: 5},
		{Id: "s6", Name: "student6", Age: 6},
		{Id: "s7", Name: "student7", Age: 7},
		{Id: "s8", Name: "student8", Age: 8},
		{Id: "s9", Name: "student9", Age: 9},
	}

	studentIds := make([]string, 0, len(students))
	for _, student := range students {
		studentIds = append(studentIds, student.Id)
	}

	SetStudents(studentServiceClient, students)
	GetStudents(studentServiceClient, studentIds)
}

func SetStudents(studentServiceClient studentpb.StudentServiceClient, students []*studentpb.Student) {
	log.Println("<----------------------SetStudents---------------------->")
	defer log.Println("<----------------------SetStudents---------------------->")

	for _, student := range students {
		if _, err := studentServiceClient.SetStudent(context.Background(), student); err != nil {
			log.Fatalf("error setting student: %v\n", err)
		}
		log.Printf("successfully set student: %v\n", student.Name)
	}
}

func GetStudents(studentServiceClient studentpb.StudentServiceClient, studentIds []string) {
	log.Println("<----------------------GetStudents---------------------->")
	defer log.Println("<----------------------GetStudents---------------------->")
	var student *studentpb.Student
	var err error

	for _, studentId := range studentIds {
		req := studentpb.GetStudentRequest{Id: studentId}
		if student, err = studentServiceClient.GetStudent(context.Background(), &req); err != nil {
			log.Fatalf("error getting student %s: %v\n", studentId, err)
		}
		log.Printf("student: %v\n", student)
	}
}
