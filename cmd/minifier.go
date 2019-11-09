package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"

	minifyv1 "github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	minifyv2 "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

type hashedInfo struct {
	extension string
	hashed    string
	minified  []byte
}

func minify(fileName string, fileBytes []byte) (*hashedInfo, error) {
	emptyM := map[string]string{}
	m1 := minifyv1.New()
	m2 := minifyv2.New()
	var extension string
	var minifiedBytes []byte
	fileBytesBuf := bytes.NewBuffer(fileBytes)
	if strings.HasSuffix(fileName, ".css") {
		var out bytes.Buffer
		err := css.Minify(m2, &out, fileBytesBuf, emptyM)
		if err != nil {
			return nil, err
		}
		extension = "css"
		minifiedBytes = out.Bytes()
	} else if strings.HasSuffix(fileName, ".js") {
		var out bytes.Buffer
		err := js.Minify(m1, &out, fileBytesBuf, emptyM)
		if err != nil {
			return nil, err
		}
		extension = "js"
		minifiedBytes = out.Bytes()
	}
	fileHash := fmt.Sprintf("%x", md5.Sum(minifiedBytes))
	return &hashedInfo{
		extension: extension,
		hashed:    fileHash,
		minified:  minifiedBytes,
	}, nil
}
