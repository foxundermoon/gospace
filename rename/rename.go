package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	reges := regexp.MustCompile(`(.+)([^\.]+$)`)
	allFiles := make(map[string]string)
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		allFiles[f.Name()] = reges.ReplaceAllString(f.Name(), "$1")
		return nil
	})
	fmt.Println(allFiles)

}
