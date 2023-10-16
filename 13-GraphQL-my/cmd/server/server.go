package main

import (
	"database/sql"
	"github.com/esirangelomub/devfullcycle-goexpert/13-GraphQL-my/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/esirangelomub/devfullcycle-goexpert/13-GraphQL-my/graph"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Read SQL script from file
	sqlScript, err := os.ReadFile("./db.sql")
	if err != nil {
		log.Fatalf("could not read SQL file: %v", err)
	}

	// Execute the SQL script
	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatalf("could not initialize db: %v", err)
	}

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)
	lessonDB := database.NewLesson(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
		LessonDB:   lessonDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
