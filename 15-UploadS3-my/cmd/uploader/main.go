package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(
			"AKID",
			"1234",
			"",
		),
	})
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "go-expert-bucket-example"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
}

func uploadFile(filename string) {
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Start to upload file %s to bucket %s\n", completeFileName, s3Bucket)

	file, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error to open file %s to bucket %s\n", completeFileName, s3Bucket)
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error to upload file %s to bucket %s\n", completeFileName, s3Bucket)
		return
	}
	fmt.Printf("Success to upload file %s to bucket %s\n", completeFileName, s3Bucket)

}
