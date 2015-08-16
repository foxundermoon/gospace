package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//reges := regexp.MustCompile(`(.+)([^\.]+$)`)
	allFiles := make(map[string]string)
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		//allFiles[f.Name()] = reges.ReplaceAllString(f.Name(), "$1")
		filename := strings.TrimRight(f.Name(), filepath.Ext(f.Name()))
		ext := filepath.Ext(filename)
		if (Mstr{ext}).Contains("tar|zip|gz") {
			filename = strings.TrimRight(filename, ext)
		}
		allFiles[f.Name()] = filename
		return nil
	})
	fmt.Println(allFiles)
	if err != nil {
		fmt.Println(err)
	}

}

type Mstr struct {
	string
}

func (s Mstr) Contains(tail string) bool {
	fmt.Println(s.string)
	var tails []string
	if strings.Contains(tail, "|") {
		tails = strings.Split(tail, "|")
	} else {
		tails = []string{tail}
	}
	for _, v := range tails {
		if strings.Contains(s.string, v) {
			return true
		}
	}
	return false
}
