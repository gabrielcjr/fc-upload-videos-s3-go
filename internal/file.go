package internal

import (
	"fmt"
	"log"
	"os"

	vidio "github.com/AlexEidt/Vidio"
)

type Videos struct {
	Repo    string
	Chapter string
}

func (v *Videos) CreateFileVideosDuration() {
	if v.Repo == "" || v.Chapter == "" {
		fmt.Println("Error: repo and chapter are required")
		os.Exit(1)
	}
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
	if fullPath == true {
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
}

func (v *Videos) SaveInFile(fileName []string) {
	f, err := os.Create("videos.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, fileName := range fileName {
		fileName += "\n"
		_, err2 := f.WriteString(fileName)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Printf("Video %s added\n", fileName)
}

func (v *Videos) GetDuration(pathToVideo string) float64 {
	video, err := vidio.NewVideo(pathToVideo)
	if err != nil {
		panic(err)
	}
	return video.Duration()
}

func (v *Videos) FormatTime(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}

func (v *Videos) RemoveAccents(s string) string {
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

func (v *Videos) RenameFiles() {
	files, err := v.GetFilesPath(true)
	// fmt.Print(files)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
		fmt.Println(v.RemoveAccents(file))
		os.Rename(file, v.RemoveAccents(file))
	}
}
