package graph

import "github.com/esirangelomub/devfullcycle-goexpert/13-GraphQL-my/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
	LessonDB   *database.Lesson
}
