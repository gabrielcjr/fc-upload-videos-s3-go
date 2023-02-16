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

func AwsClient(region string, id string, secretKey string) *s3.Client {
	godotenv.Load()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(id, secretKey, "")),
	)
	if err != nil {
		log.Fatal(err)
	}
	return s3.NewFromConfig(cfg)
}

func (a *AWSUpload) UploadVideos(wg *sync.WaitGroup, cl *s3.Client) {
	file, err := os.Open(a.VideosLocalPath)

	pathToS3 := a.S3Repo + a.S3Chapter + a.FileName

	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", a.FileName, err)
	} else {
		defer file.Close()
		_, err := cl.PutObject(context.TODO(), &s3.PutObjectInput{
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

func (a *AWSUpload) ChangePathToPublicRead(cl *s3.Client) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_READ")),
		Prefix: aws.String(a.S3Repo + a.S3Chapter + a.FileName),
	}
	contents, err := cl.ListObjectsV2(context.Background(), params)

	if err != nil {
		panic(err)
	}

	result := contents.Contents

	wg := sync.WaitGroup{}
	wg.Add(len(result))

	for _, content := range result {
		fmt.Println(*content.Key)
		go func() {
			defer wg.Done()
			_, err := cl.PutObjectAcl(context.TODO(), &s3.PutObjectAclInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_READ")),
				Key:    content.Key,
				AccessControlPolicy: &types.AccessControlPolicy{
					Grants: []types.Grant{
						{
							Grantee: &types.Grantee{
								Type:        "CanonicalUser",
								DisplayName: aws.String(os.Getenv("AWS_DISPLAY_NAME")),
								ID:          aws.String(os.Getenv("AWS_OWNER_ID")),
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
						DisplayName: aws.String(os.Getenv("AWS_DISPLAY_NAME")),
						ID:          aws.String(os.Getenv("AWS_OWNER_ID")),
					},
				},
			})
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
}
