package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/Edigiraldo/grpc/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (student *models.Student, err error) {
	student = &models.Student{}
	row := repo.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err = row.Scan(&student.Id, &student.Name, &student.Age); err != nil {
		return nil, err
	}

	return student, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1 ,$2, $3)",
		student.Id, student.Name, student.Age)

	return err
}

func (repo *PostgresRepository) GetExam(ctx context.Context, id string) (exam *models.Exam, err error) {
	exam = &models.Exam{}
	row := repo.db.QueryRowContext(ctx, "SELECT id, name FROM exams WHERE id = $1", id)
	if err = row.Scan(&exam.Id, &exam.Name); err != nil {
		return nil, err
	}

	return exam, nil
}

func (repo *PostgresRepository) SetExam(ctx context.Context, exam *models.Exam) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO exams (id, name) VALUES ($1 ,$2)",
		exam.Id, exam.Name)

	return err
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, question, answer, exam_id) VALUES ($1 ,$2, $3, $4)",
		question.Id, question.Question, question.Answer, question.ExamId)

	return err
}

func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enroll *models.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollments (student_id, exam_id) VALUES ($1 ,$2)",
		enroll.StudentId, enroll.ExamId)

	return err
}

func (repo *PostgresRepository) GetStudentsPerExam(ctx context.Context, examId string) ([]models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students RIGHT JOIN (SELECT student_id FROM enrollments WHERE exam_id = $1) as students_on_exam ON students.id = students_on_exam.student_id", examId)

	if err != nil {
		return nil, err
	}

	var students []models.Student
	for rows.Next() {
		student := models.Student{}
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, student)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rows.Close()
	if err != nil {
		log.Printf("error closing rows: %v", err)
		return nil, err
	}

	return students, nil
}

func (repo *PostgresRepository) GetQuestionsPerExam(ctx context.Context, examId string) ([]models.Question, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, question, exam_id FROM questions WHERE exam_id = $1", examId)
	if err != nil {
		return nil, err
	}

	var questions []models.Question
	for rows.Next() {
		question := models.Question{}
		if err := rows.Scan(&question.Id, &question.Question, &question.ExamId); err == nil {
			questions = append(questions, question)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rows.Close()
	if err != nil {
		log.Printf("error closing rows: %v", err)
		return nil, err
	}

	return questions, nil
}
