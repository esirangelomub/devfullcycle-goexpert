package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description *string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) "+
		"VALUES ($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Course{ID: id, Name: name, Description: &description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) FindOne(id string) (Course, error) {
	row := c.db.QueryRow("SELECT id, name, description, category_id FROM courses WHERE id = $1", id)
	var course Course
	err := row.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) FindByLessonID(lessonID string) (Course, error) {
	row := c.db.QueryRow("SELECT co.id, co.name, co.description FROM courses co "+
		"INNER JOIN lessons le ON co.id = le.course_id WHERE le.id = $1", lessonID)
	var course Course
	err := row.Scan(&course.ID, &course.Name, &course.Description)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}
