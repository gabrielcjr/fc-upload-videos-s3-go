package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var videos = Videos{
	Repo:    "1",
	Chapter: "1",
}

func TestGetFilesPath(t *testing.T) {
	err := os.Mkdir("videos", os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}

	f, _ := os.Create("./videos/test")

	defer f.Close()

	defer os.RemoveAll("./videos")

	expected, _ := os.Getwd()

	expected += "/videos/test"

	result, err := videos.GetFilesPath(true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	assert.Equal(t, expected, result[0])

	expected = "test"

	result, err = videos.GetFilesPath(false)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, expected, result[0])
}

func TestSaveInFile(t *testing.T) {
	fileName := []string{"videos1.mp4"}

	videos.saveInFile(fileName)

	defer os.Remove("videos.md")

	result, _ := filepath.Glob("videos.md")

	expected := "videos.md"

	assert.Equal(t, expected, result[0])

	result1, _ := os.ReadFile("videos.md")

	expected = "videos1.mp4"

	assert.Equal(t, expected, string(result1))
}

func TestGetDuration(t *testing.T) {
	out, _ := os.Create("videoTest.mp4")
	fileToDownload := "https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4"

	resp, _ := http.Get(fileToDownload)
	defer resp.Body.Close()
	n, _ := io.Copy(out, resp.Body)

	fmt.Println(n)

	defer os.Remove("videoTest.mp4")

	pathToFile := "videoTest.mp4"

	duration := videos.getDuration(pathToFile)

	assert.Equal(t, duration, 14)
}

func FuzzFormatTime(f *testing.F) {
	seed := []int{-10, 1, 10, 180, 203, 360, 600, 2000, 7099}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount int) {
		result, err := videos.formatTime(amount)
		fmt.Println(result)
		fmt.Println(err)
		if err != nil {
			expected := errors.New("Use um valor maior que 0 e menor que 5999")
			assert.Equal(t, expected, err)
		} else {
			minutes := amount / 60
			seconds := amount % 60
			expected := fmt.Sprintf("%02d:%02d", minutes, seconds)
			assert.Equal(t, expected, result)
		}
	})
}

func TestNormalizeFilename(t *testing.T) {

	input := []string{"áé ilksjdb", "klkjf 234 áéç", "sdfbIYV", "áç dfb"}

	expected := []string{"ae-ilksjdb", "klkjf-234-aec", "sdfbIYV", "ac-dfb"}

	for k, value := range input {
		result := videos.normalizeFilename(value)
		assert.Equal(t, result, expected[k])
	}
}

func TestRenameFiles(t *testing.T) {
	_ = os.Mkdir("videos", os.ModePerm)

	f, err := os.Create("./videos/filé.mp4")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	defer os.RemoveAll("./videos")

	videos.renameFiles()

	expected, _ := os.Getwd()

	expected += "/videos/file.mp4"

	result, _ := filepath.Glob(expected)

	assert.Equal(t, expected, result[0])

}

// func FuzzNormalizeFilename(f *testing.F) {
// 	seed := []string{"wrfgsg", "áéí", "ibjawrf89734", "edrfg sdfbsdfb", "çççttt123454  áá sdfvb"}

// 	for _, fileName := range seed {
// 		f.Add(fileName)
// 	}

// 	f.Fuzz(func(t *testing.T, fileName string) {
// 		result := videos.normalizeFilename(fileName)
// 		m1 := regexp.MustCompile("/[^a-zA-Z0-9 ]/g")
// 		expected := m1.ReplaceAllString(fileName, "${1}.${2}")

// 		assert.Equal(t, expected, result)

// 	})

// }
