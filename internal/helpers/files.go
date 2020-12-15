package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func StrReplaceInDirectory(dir, str, replacement string) error {
	filesList, err := GetFilesPathsList(dir)
	if err != nil {
		return err
	}

	for i := range filesList {
		if err := StrReplaceInFile(filesList[i], str, replacement); err != nil {
			return err
		}
	}

	return nil
}

func StrReplaceInFile(filepath, str, replacement string) error {
	read, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	newContents := strings.Replace(string(read), str, replacement, -1)

	return ioutil.WriteFile(filepath, []byte(newContents), 0)
}

func GetFilesPathsList(dir string) ([]string, error) {
	var filesList []string

	err := filepath.Walk(dir, func(curPath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("folder traversing error: %s", err)
		}

		if !info.IsDir() {
			relFilePath := strings.ReplaceAll(curPath, dir, "")
			filesList = append(filesList, relFilePath)
		}

		return nil
	})

	return filesList, err
}
