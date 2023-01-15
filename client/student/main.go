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
		{Id: "s1", Name: "Dennis Ritchie", Age: 70},
		{Id: "s6", Name: "Brian Kernighan", Age: 81},
		{Id: "s2", Name: "Bjarne Stroustrup", Age: 72},
		{Id: "s3", Name: "Linus Torvalds", Age: 53},
		{Id: "s4", Name: "Anders Hejlsberg", Age: 62},
		{Id: "s9", Name: "Donald Knuth", Age: 85},
		{Id: "s8", Name: "Guido van Rossum", Age: 66},
		{Id: "s5", Name: "Tim Berners-Lee", Age: 67},
		{Id: "s7", Name: "Ken Thompson", Age: 79},
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
