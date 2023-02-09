package main

import (
	"fmt"
	"math"

	"github.com/gacarneirojr/fc-upload-videos-s3-go/internal"
)

func main() {

	file := internal.Videos{
		Repo:    "1",
		Chapter: "2",
	}
	// Testando file.GetFilesPath

	result, err := file.GetFilesPath(false)
	if err != nil {
		panic(err)
	}

	file.SaveInFile(result)

	videoTime := file.GetDuration("./videos/22.01-boás-vindas-ao-modulo-de-infrastructure-de-video.mp4")

	roundVideoTime := int(math.Round(videoTime))

	result1 := file.FormatTime(roundVideoTime)

	fmt.Print((result1))

	result2 := file.RemoveAccents("áéíóú ")

	fmt.Print(result2)

	file.RenameFiles()

}
