package main

import (
	"fmt"

	"github.com/gacarneirojr/fc-upload-videos-s3-go/internal"
)

func main() {

	file := internal.Videos{
		Repo:    "1",
		Chapter: "2",
	}

	result, err := file.GetFilesPath("Não")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}
