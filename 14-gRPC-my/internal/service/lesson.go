package service

import (
	"context"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/database"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type LessonService struct {
	pb.UnimplementedLessonServiceServer
	lessonDB database.Lesson
}

func NewLessonService(lessonDB database.Lesson) *LessonService {
	return &LessonService{
		lessonDB: lessonDB,
	}
}

func (l *LessonService) CreateLesson(ctx context.Context, in *pb.CreateLessonRequest) (*pb.Lesson, error) {
	lesson, err := l.lessonDB.Create(in.Name, in.Description, in.Content, in.CourseId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating lesson")
	}

	lessonResponse := &pb.Lesson{
		Id:          lesson.ID,
		Name:        lesson.Name,
		Description: *lesson.Description,
		Content:     *lesson.Content,
	}

	return lessonResponse, nil
}

func (l *LessonService) ListLessons(context.Context, *pb.Blank) (*pb.LessonList, error) {
	lessons, err := l.lessonDB.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing lessons")
	}

	var lessonsResponse []*pb.Lesson

	for _, course := range lessons {
		lessonResponse := &pb.Lesson{
			Id:          course.ID,
			Name:        course.Name,
			Description: *course.Description,
			Content:     *course.Content,
		}

		lessonsResponse = append(lessonsResponse, lessonResponse)
	}

	return &pb.LessonList{Lessons: lessonsResponse}, nil
}

func (l *LessonService) GetLesson(ctx context.Context, in *pb.LessonGetRequest) (*pb.Lesson, error) {
	lesson, err := l.lessonDB.FindOne(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting lesson")
	}

	lessonResponse := &pb.Lesson{
		Id:          lesson.ID,
		Name:        lesson.Name,
		Description: *lesson.Description,
		Content:     *lesson.Content,
	}

	return lessonResponse, nil
}

func (l *LessonService) CreateLessonStream(stream pb.LessonService_CreateLessonStreamServer) error {
	lessons := &pb.LessonList{}

	for {
		lesson, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(lessons)
		}

		if err != nil {
			return err
		}

		lessonResponse, err := l.CreateLesson(context.Background(), lesson)
		if err != nil {
			return err
		}

		lessons.Lessons = append(lessons.Lessons, &pb.Lesson{
			Id:          lessonResponse.Id,
			Name:        lessonResponse.Name,
			Description: lessonResponse.Description,
			Content:     lessonResponse.Content,
		})
	}
}

func (l *LessonService) CreateLessonStreamBidirectional(stream pb.LessonService_CreateLessonStreamBidirectionalServer) error {
	for {
		lesson, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		lessonResponse, err := l.CreateLesson(context.Background(), lesson)
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.Lesson{
			Id:          lessonResponse.Id,
			Name:        lessonResponse.Name,
			Description: lessonResponse.Description,
			Content:     lessonResponse.Content,
		}); err != nil {
			return err
		}
	}
}
