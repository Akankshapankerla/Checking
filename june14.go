package main

import (
	"fmt"
	"io/ioutil"
)

type FileError struct {
	Path string
	Op   string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error: %s %s: %v", e.Op, e.Path, e.Err)
}

func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadFileWithCustomError(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", &FileError{
			Path: filePath,
			Op:   "read",
			Err:  err,
		}
	}
	return string(data), nil
}

func main14() {
	filePath := "example.txt"

	content, err := ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s\n", content)
}

func mainWithCustomError() {
	filePath := "example.txt"

	content, err := ReadFileWithCustomError(filePath)
	if err != nil {
		if fileErr, ok := err.(*FileError); ok {
			fmt.Printf("Error reading file '%s' during operation '%s': %v\n", fileErr.Path, fileErr.Op, fileErr.Err)
		} else {
			fmt.Printf("Error reading file: %v\n", err)
		}
		return
	}

	fmt.Printf("File content:\n%s\n", content)
}
