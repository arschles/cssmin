package main

import (
	"io"
	"os"
)

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
//
// Thank you to https://golangcode.com/writing-to-file/ for this code
func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
