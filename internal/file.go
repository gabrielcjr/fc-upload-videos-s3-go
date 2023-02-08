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

// func (v *Videos) filenameNormalized(file string) {
// 	characters := map[string]string{
// 		"a": "á|à|ã|â",
// 		"A": "Á|À|Ã|Â",
// 		"e": "é|ê",
// 		"E": "É|Ê",
// 		"i": "í",
// 		"I": "Í",
// 		"o": "ó|õ|ô",
// 		"O": "Ó|Õ|Ô",
// 		"u": "ú",
// 		"U": "Ú",
// 		"c": "ç",
// 		"C": "Ç",
// 		"-": " ",
// 	  }

// 	  str := file
// 	  for k, name := range characters {
// 		str = str.replace(new RegExp)
// 		}
// 	}
// 	  for (const i in map) {
// 		str = str.replace(new RegExp(map[i], 'g'), i)
// 	  }
// 	  return str
// }

func (v *Videos) renameFiles() {
	// To implement
}
