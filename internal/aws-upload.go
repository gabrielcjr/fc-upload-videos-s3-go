package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type AWSUpload struct {
	S3Repo          string
	S3Chapter       string
	FileName        string
	VideosLocalPath string
}

func (a *AWSUpload) UploadVideos(wg *sync.WaitGroup) {
	err := godotenv.Load()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(a.VideosLocalPath)

	pathToS3 := a.S3Repo + a.S3Chapter + a.FileName

	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", a.FileName, err)
	} else {
		defer file.Close()
		_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_UPLOAD")),
			Key:    aws.String(pathToS3),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				a.FileName, os.Getenv("AWS_BUCKET_NAME_UPLOAD"), pathToS3, err)
		}
		fmt.Println("File uploaded to", pathToS3)
		wg.Done()
	}
}
