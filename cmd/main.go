package main

import (
	"github.com/gacarneirojr/fc-upload-videos-s3-go/internal"
)

func main() {

	file := internal.Videos{
		Repo:    "golang/",
		Chapter: "2/",
	}

	file.CreateFileVideosDuration()

}
