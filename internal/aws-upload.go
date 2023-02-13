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
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/joho/godotenv"
)

type AWSUpload struct {
	S3Repo          string
	S3Chapter       string
	FileName        string
	VideosLocalPath string
}

func (a *AWSUpload) Client(region string, id string, secretKey string) {
	// err := godotenv.Load()
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region),
	// 	config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(id, secretKey, "")),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return s3.NewFromConfig(cfg)
}

func (a *AWSUpload) UploadVideos(wg *sync.WaitGroup) {
	// client := a.Client(
	// 	os.Getenv("AWS_REGION"),
	// 	os.Getenv("AWS_ACCESS_KEY_ID"),
	// 	os.Getenv("AWS_SECRET_ACCESS_KEY"))
	err := godotenv.Load()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")),
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

func (a *AWSUpload) ChangePathToPublicRead() {
	// client := a.Client("us-east-1", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	err := godotenv.Load()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_READ")),
		Prefix: aws.String(a.S3Repo + a.S3Chapter + a.FileName),
	}

	contents, err := client.ListObjectsV2(context.Background(), params)

	if err != nil {
		panic(err)
	}

	result := contents.Contents

	for _, content := range result {
		fmt.Println(*content.Key)
		_, err := client.PutObjectAcl(context.TODO(), &s3.PutObjectAclInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_READ")),
			Key:    content.Key,
			AccessControlPolicy: &types.AccessControlPolicy{
				Grants: []types.Grant{
					{
						Grantee: &types.Grantee{
							Type:        "CanonicalUser",
							DisplayName: aws.String("wesleywillians"),
							ID:          aws.String("a3edb89dc8762b1d543412e1b0999c8b17e8a1e94c3694bf2e35d4b61499419d"),
						}, Permission: "FULL_CONTROL",
					},
					{
						Grantee: &types.Grantee{
							Type: "Group",
							URI:  aws.String("http://acs.amazonaws.com/groups/global/AllUsers"),
						}, Permission: "READ",
					},
				},
				Owner: &types.Owner{
					DisplayName: aws.String("wesleywillians"),
					ID:          aws.String("a3edb89dc8762b1d543412e1b0999c8b17e8a1e94c3694bf2e35d4b61499419d"),
				},
			},
		})
		if err != nil {
			panic(err)
		}
	}
}
