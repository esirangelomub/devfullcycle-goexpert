package service

import (
	"context"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/database"
	"github.com/esirangelomub/devfullcycle-goexpert/14-gRPC-my/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.Course
}

func NewCourseService(courseDB database.Course) *CourseService {
	return &CourseService{
		CourseDB: courseDB,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.Course, error) {
	course, err := c.CourseDB.Create(in.Name, in.Description, in.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating course")
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: *course.Description,
	}

	return courseResponse, nil
}

func (c *CourseService) ListCourses(ctx context.Context, in *pb.Blank) (*pb.CourseList, error) {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing courses")
	}

	var coursesResponse []*pb.Course

	for _, course := range courses {
		courseResponse := &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: *course.Description,
		}

		coursesResponse = append(coursesResponse, courseResponse)
	}

	return &pb.CourseList{Courses: coursesResponse}, nil
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.CourseGetRequest) (*pb.Course, error) {
	course, err := c.CourseDB.FindOne(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting course")
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: *course.Description,
	}

	return courseResponse, nil
}

func (c *CourseService) CreateCourseStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CourseList{}

	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}

		if err != nil {
			return err
		}

		courseResponse, err := c.CreateCourse(context.Background(), course)
		if err != nil {
			return err
		}

		courses.Courses = append(courses.Courses, &pb.Course{
			Id:          courseResponse.Id,
			Name:        courseResponse.Name,
			Description: courseResponse.Description,
		})
	}
}

func (c *CourseService) CreateCourseStreamBidirectional(stream pb.CourseService_CreateCourseStreamBidirectionalServer) error {
	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		courseResponse, err := c.CreateCourse(context.Background(), course)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Course{
			Id:          courseResponse.Id,
			Name:        courseResponse.Name,
			Description: courseResponse.Description,
		})
		if err != nil {
			return err
		}
	}
}
