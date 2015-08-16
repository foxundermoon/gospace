package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	r, err := sha256Sum("c:\\dev\\vim\\vim74\\gvim.exe")
	if err != nil {
		fmt.Println(err)
	}

	log.Println(r)
	fmt.Println("hello")

}

func sha256Sum(path string) (result string, err error) {
	fp, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	file, err := os.Open(fp)
	defer file.Close()
	if err != nil {
		return "", err
	}
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil

}
