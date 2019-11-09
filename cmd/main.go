package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/markbates/pkger"
)

func main() {
	log.Printf("%s", os.Args)
	if len(os.Args) != 2 {
		log.Fatalf("Pass the directory name here please!")
	}
	dir := os.Args[1]
	pkgerDir := fmt.Sprintf("/%s", dir)
	// a mapping between raw file names and hashed file names
	fileMap := map[string]string{}
	pkger.Walk(
		pkgerDir,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			f, err := pkger.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			fileBytes, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}
			hashedInfo, err := minify(path, fileBytes)
			if err != nil {
				return err
			}
			newFileName := fmt.Sprintf(
				"%s.%s.min.%s",
				path,
				hashedInfo.hashed,
				hashedInfo.extension,
			)
			f, err = pkger.Create(newFileName)
			if err != nil {
				return err
			}
			if _, err := f.Write(hashedInfo.minified); err != nil {
				return err
			}
			fileMap[path] = newFileName
			return nil
		},
	)
	for origFile, hashedFile := range fileMap {
		log.Printf(
			"%s\n\t ==> Becomes <==\n%s",
			origFile,
			hashedFile,
		)
	}
}
