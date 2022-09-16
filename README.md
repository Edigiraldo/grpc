To start your postgress db into a Docker container go to database directory and run:
- docker build . -t grpc-db
- docker run -p 54321:5432 grpc-db

Then, go to the grpc dir and start the student and exam servers:
- go run cmd/student/main.go
- go run cmd/exam/main.go

Now the student service will allow you to set and get students. Also, in the exam service, you will be able to set exams, the questions for them, enroll students to each exam, get the studens per exam and finally take an exam. To use these services:
- Go to Postman and create a new gRPC request.
- Connect to localhost:5060(Students) or localhost:5070(Exams).
- Click on select a method, select the method you want, for example SetExam.
- Click on the "Generate Example Message" fill the generated message and invoke the procedure.

If you want to run an already prepared example:
- go run client/student/main.go
- go run client/exam/main.go
