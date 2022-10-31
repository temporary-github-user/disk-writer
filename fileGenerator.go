package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func fileGenerator(id, size int, rootPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(filepath.Join(rootPath, filepath.Base(strconv.Itoa(id)+"_generated")))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.CopyN(file, rand.Reader, MB*int64(size))
	if err != nil {
		log.Fatal(err)
	}
}
