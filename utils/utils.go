package utils

import (
	"os"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func SlicePath(path string, target string) string {
	files := strings.Split(path, "/")
	var result []string
	for i, file := range files {
		if file == target {
			result =  files[i + 1:]
			break
		}
	}
	return strings.Join(result, "/")
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}