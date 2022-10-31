package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
)

func TestFileGenerator(t *testing.T) {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fileGenerator(i, 5, ".", wg)
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		file := filepath.Join(".", filepath.Base(strconv.Itoa(i)+"_generated"))
		if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
			t.Error("file not found:", file)
		}
		err := os.Remove(file)
		if err != nil {
			log.Fatal(err)

		}
	}

}
