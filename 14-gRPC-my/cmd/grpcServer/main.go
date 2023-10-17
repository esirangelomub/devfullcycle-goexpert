package main

import (
	"database/sql"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/database"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/pb"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

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

	// Create a new CategoryDB
	categoryDB := database.NewCategory(db)
	// Create a new CourseDB
	courseDB := database.NewCourse(db)
	// Create a new LessonDB
	lessonDB := database.NewLesson(db)

	// Create a new CategoryService
	categoryService := service.NewCategoryService(*categoryDB)
	// Create a new CourseService
	courseService := service.NewCourseService(*courseDB)
	// Create a new LessonService
	lessonService := service.NewLessonService(*lessonDB)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the CategoryService
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	// Register the CourseService
	pb.RegisterCourseServiceServer(grpcServer, courseService)
	// Register the LessonService
	pb.RegisterLessonServiceServer(grpcServer, lessonService)

	// Register the reflection service
	reflection.Register(grpcServer)

	// Open a TCP connection at port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
