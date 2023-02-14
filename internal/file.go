package internal

import (
	"fmt"
	"log"
	"math"
	"os"

	vidio "github.com/AlexEidt/Vidio"
)

type Videos struct {
	Repo    string
	Chapter string
}

func (v *Videos) CreateFileVideosDuration() {
	if v.Repo == "" || v.Chapter == "" {
		fmt.Println("Error: repo and/or chapter are required")
		os.Exit(1)
	}
	v.renameFiles()

	files, err := v.GetFilesPath(false)

	if err != nil {
		panic(err)
	}

	var videos []string

	for _, file := range files {
		seconds := v.getDuration("./videos/" + file)
		line := v.Repo + v.Chapter + file + " " + v.formatTime(seconds) + "\r\n"
		videos = append(videos, line)
	}

	v.saveInFile(videos)

	fmt.Println("File created successfully")
}

func (v *Videos) GetFilesPath(fullPath bool) ([]string, error) {
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
	if fullPath {
		for _, fileName := range files {
			fileNameString := fileName.Name()
			fullPath := path + "/videos/" + fileNameString
			list = append(list, fullPath)
		}
		return list, nil
	}

	for _, fileName := range files {
		fileNameString := fileName.Name()
		list = append(list, fileNameString)
	}
	return list, nil
}

func (v *Videos) saveInFile(fileName []string) {
	f, err := os.Create("videos.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, fileName := range fileName {
		_, err2 := f.WriteString(fileName)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func (v *Videos) getDuration(pathToVideo string) int {
	video, err := vidio.NewVideo(pathToVideo)
	if err != nil {
		panic(err)
	}

	rawTime := video.Duration()

	return int(math.Round(rawTime))
}

func (v *Videos) formatTime(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}

func (v *Videos) normalizeFilename(s string) string {
	// A map of accented characters and their equivalent without accents
	accents := map[rune]rune{
		'À': 'A', 'Á': 'A', 'Â': 'A', 'Ã': 'A', 'Ä': 'A', 'Å': 'A',
		'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a', 'å': 'a',
		'È': 'E', 'É': 'E', 'Ê': 'E', 'Ë': 'E',
		'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e',
		'Ì': 'I', 'Í': 'I', 'Î': 'I', 'Ï': 'I',
		'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i',
		'Ò': 'O', 'Ó': 'O', 'Ô': 'O', 'Õ': 'O', 'Ö': 'O', 'Ø': 'O',
		'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ö': 'o', 'ø': 'o',
		'Ù': 'U', 'Ú': 'U', 'Û': 'U', 'Ü': 'U',
		'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u',
		'Ý': 'Y', 'ý': 'y', 'ÿ': 'y', ' ': '-',
	}

	var result []rune
	for _, r := range s {
		if replacement, ok := accents[r]; ok {
			result = append(result, replacement)
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func (v *Videos) renameFiles() {
	files, err := v.GetFilesPath(false)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		file = "./videos/" + file
		os.Rename(file, v.normalizeFilename(file))
	}
}
