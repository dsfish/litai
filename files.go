package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const textDir string = "texts"

func listFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var allPaths []string
	for _, file := range files {
		if file.IsDir() {
			paths, err := listFiles(fmt.Sprintf("%s/%s", dir, file.Name()))
			if err != nil {
				return nil, err
			}
			allPaths = append(allPaths, paths...)
		} else if !strings.HasPrefix(file.Name(), ".") { // ignore hidden files
			allPaths = append(allPaths, fmt.Sprintf("%s/%s", dir, file.Name()))
		}
	}

	return allPaths, nil
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}