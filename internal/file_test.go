package internal

import (
	"fmt"
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
	f, _ := os.Create("./videos/test")

	defer f.Close()

	defer os.Remove("./videos/test")

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
	fileName := []string{"videos.md"}

	videos.saveInFile(fileName)

	defer os.Remove("videos.md")

	result, _ := filepath.Glob("videos.md")

	expected := "videos.md"

	assert.Equal(t, expected, result[0])
}
