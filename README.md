- docker build . -t grpc-db
- docker run -p 54321:5432 grpc-db

- go run cmd/student/main.go
- go run cmd/exam/main.go
- Go to Postman and create a new gRPC request. Connect to localhost:5060(Students) or localhost:5070(Exams) Click on select a method, select the method, click on the "Generate Example Message" fill the message and invoke the procedure.
