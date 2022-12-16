package main

import (
	"fmt"
	"os"
)

type File struct {
	repo    string
	chapter string
}

func (f *File) createFileVideosDuration() {
	if f.repo == "" || f.chapter == "" {
		fmt.Println("Error: repo and chapter are required")
		os.Exit(1)
	}
}

func (f *File) getFilesPath(completePath string) {
	// To implement
}

func (f *File) saveInFile(videos []string) {
	// To implement
}

func (f *File) fromatTime(seconds string) {
	// To implement
}

func (f *File) filenameNormalized(file string) {
	// To implement
}

func (f *File) renameFiles() {
	// To implement
}
