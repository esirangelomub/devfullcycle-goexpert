package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Lesson struct {
	db          *sql.DB
	ID          string
	Name        string
	Description *string
	Content     *string
	CourseID    string
}

func NewLesson(db *sql.DB) *Lesson {
	return &Lesson{db: db}
}

func (l *Lesson) Create(name string, description string, courseID string, content string) (*Lesson, error) {
	id := uuid.New().String()
	_, err := l.db.Exec("INSERT INTO lessons (id, name, description, content, course_id) "+
		"VALUES ($1, $2, $3, $4, $5)", id, name, description, content, courseID)
	if err != nil {
		return nil, err
	}
	return &Lesson{ID: id, Name: name, Description: &description, Content: &content, CourseID: courseID}, nil
}

func (l *Lesson) FindAll() ([]Lesson, error) {
	rows, err := l.db.Query("SELECT id, name, description, content, course_id FROM lessons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.Content, &lesson.CourseID)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func (l *Lesson) FindByCourseID(courseID string) ([]Lesson, error) {
	rows, err := l.db.Query("SELECT id, name, description, content, course_id FROM lessons WHERE course_id = $1", courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.Content, &lesson.CourseID)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
