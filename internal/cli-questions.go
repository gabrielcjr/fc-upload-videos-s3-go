package internal

import "github.com/manifoldco/promptui"

var Repositories = map[string]string{
	"MBA":           "code/mba/solution-architecture/",
	"TYPESCRIPT":    "code/fullcycle/fc3/microsservico-catalogo-de-videos-com-typescript/",
	"DOTNET":        "code/fullcycle/fc3/microsservico-catalogo-de-videos-com-dotnet/",
	"REACT":         "code/fullcycle/fc3/microsservico-administracao-do-catalogo-de-videos-com-React/",
	"JAVA":          "code/fullcycle/fc3/microsservico-catalogo-de-videos-com-java-new/",
	"PHP":           "code/fullcycle/fc3/microsservico-catalogo-de-videos-com-php/",
	"PYTHON":        "code/fullcycle/fc3/microsservico-catalogo-de-videos-com-python/",
	"DEPLOY_CLOUDS": "code/fullcycle/Deploy-das-Cloud-Providers/",
	"GOLANG":        "code/go/",
	"EDA":           "code/fullcycle/fc3/EDA/",
}

var (
	Prompt = promptui.Select{
		Label: "Selecione o repositório da lista abaixo:",
		Items: []string{"MBA", "TYPESCRIPT", "DOTNET", "REACT",
			"JAVA", "PHP", "PYTHON", "DEPLOY_CLOUDS",
			"GOLANG", "EDA"},
	}
	Prompt2 = promptui.Prompt{
		Label: "Agora digite o nome/número do capítulo",
	}

	Prompt3 = promptui.Prompt{
		Label:     "Fazer upload agora?",
		IsConfirm: true,
	}

	Prompt4 = promptui.Prompt{
		Label:     "Deseja alterar as permissões no S3 agora?",
		IsConfirm: true,
	}
)
