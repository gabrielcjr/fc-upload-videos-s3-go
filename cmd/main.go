package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/gacarneirojr/fc-upload-videos-s3-go/internal"
	"github.com/joho/godotenv"
)

func main() {

	_, repo, err := internal.Prompt.Run()

	if err != nil {
		fmt.Printf("Opção incorreta %s\n", err)
	}

	chapter, err := internal.Prompt2.Run()

	if err != nil {
		fmt.Printf("Opção incorreta %s\n", err)
	}

	file := internal.Videos{
		Repo:    internal.Repositories[repo],
		Chapter: chapter + "/",
	}

	file.CreateFileVideosDuration()

	fullLocalPath, _ := file.GetFilesPath(true)
	fileNames, _ := file.GetFilesPath(false)

	isUpload, err := internal.Prompt3.Run()

	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(fileNames))

	godotenv.Load()

	client := internal.AwsClient(
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"))

	if isUpload == "y" {
		for k := range fileNames {
			aws := internal.AWSUpload{
				S3Repo:          internal.Repositories[repo],
				S3Chapter:       chapter + "/",
				FileName:        fileNames[k],
				VideosLocalPath: fullLocalPath[k],
			}
			go aws.UploadVideos(&wg, client)
		}
		wg.Wait()
	}

	isChangePermission, _ := internal.Prompt4.Run()

	if isChangePermission == "y" {
		for k := range fileNames {
			aws := internal.AWSUpload{
				S3Repo:          internal.Repositories[repo],
				S3Chapter:       chapter + "/",
				FileName:        fileNames[k],
				VideosLocalPath: fullLocalPath[k],
			}
			aws.ChangePathToPublicRead(client)
		}
	}
}
