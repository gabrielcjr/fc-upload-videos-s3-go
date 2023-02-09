package main

import (
	"fmt"

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

}
