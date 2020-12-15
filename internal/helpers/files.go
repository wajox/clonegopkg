package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

		if !info.IsDir() && !strings.Contains(curPath, ".git/") {
			filesList = append(filesList, curPath)
		}

		return nil
	})

	return filesList, err
}
func GetPkgNameFromGomod(dir string) (string, error) {
	gomodfile := fmt.Sprintf("%s/go.mod", dir)
	read, err := ioutil.ReadFile(gomodfile)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`module (.*)`)
	res := re.FindSubmatch(read)
	if len(res) < 2 {
		return "", errors.New("can not parse pkg name from go.mod")
	}

	return string(res[1]), nil
}
