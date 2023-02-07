package internal

import (
	"fmt"
	"os"
)

type Videos struct {
	Repo    string
	Chapter string
}

func (f *Videos) CreateFileVideosDuration() {
	if f.Repo == "" || f.Chapter == "" {
		fmt.Println("Error: repo and chapter are required")
		os.Exit(1)
	}
}

func (v *Videos) GetFilesPath(fullPath string) ([]string, error) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	basePath := "./videos"
	f, err := os.Open(basePath)
	if err != nil {
		panic(err)
	}

	files, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}

	var list []string
	if fullPath == "Sim" {
		for _, fileName := range files {
			fileNameString := fileName.Name()
			fullPath := path + fileNameString
			list = append(list, fullPath)
		}
		return list, nil
	}

	for _, fileName := range files {
		fileNameString := fileName.Name()
		list = append(list, fileNameString)
	}
	return list, nil

	// for _, v := range files {
	// 	fmt.Println(v.Name(), v.IsDir())
	// }

}

func (f *Videos) saveInFile(videos []string) {
	// To implement
}

func (f *Videos) fromatTime(seconds string) {
	// To implement
}

func (f *Videos) filenameNormalized(file string) {
	// To implement
}

func (f *Videos) renameFiles() {
	// To implement
}
