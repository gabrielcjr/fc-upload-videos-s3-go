package main

import (
	"fmt"
	"sync"

	"github.com/gacarneirojr/fc-upload-videos-s3-go/internal"
	"github.com/manifoldco/promptui"
)

func main() {

	prompt := promptui.Select{
		Label: "Selecione o repositório da lista abaixo:",
		Items: []string{"TYPESCRIPT", "DOTNET", "REACT",
			"JAVA", "PHP", "PYTHON", "DEPLOY_CLOUDS",
			"GOLANG", "EDA"},
	}

	_, repo, err := prompt.Run()

	if err != nil {
		fmt.Printf("Opção incorreta %s\n", err)
	}

	prompt2 := promptui.Prompt{
		Label: "Agora digite o número do capítulo",
	}

	chapter, err := prompt2.Run()

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

	prompt3 := promptui.Prompt{
		Label:     "Fazer upload agora?",
		IsConfirm: true,
	}

	isUpload, err := prompt3.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(isUpload)

	waitGroup := sync.WaitGroup{}
	fmt.Println(len(fileNames))
	waitGroup.Add(len(fileNames))

	if isUpload == "y" {
		for k, _ := range fileNames {
			aws := internal.AWSUpload{
				S3Repo:          internal.Repositories[repo],
				S3Chapter:       chapter + "/",
				FileName:        fileNames[k],
				VideosLocalPath: fullLocalPath[k],
			}
			go aws.UploadVideos(&waitGroup)
		}
		waitGroup.Wait()
	}

	prompt4 := promptui.Prompt{
		Label:     "Deseja alterar as permissões no S3 agora?",
		IsConfirm: true,
	}

	isChangePermission, err := prompt4.Run()

	if isChangePermission == "y" {
		for k, _ := range fileNames {
			aws := internal.AWSUpload{
				S3Repo:          internal.Repositories[repo],
				S3Chapter:       chapter + "/",
				FileName:        fileNames[k],
				VideosLocalPath: fullLocalPath[k],
			}
			aws.ChangePathToPublicRead()
		}
	}

}
