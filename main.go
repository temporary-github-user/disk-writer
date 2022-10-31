package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

const MB = 1024 * 1024

var locals = map[string]bool{
	"ext4":  true,
	"btrfs": true,
	"tmpfs": true,
}

type FS struct {
	Device, MountedOn, FreeSpaceInMB string
}

func main() {

	vp := viper.New()
	vp.AutomaticEnv()
	if vp.Get("LOCAL_FILESYSTEM") != nil {
		locals = map[string]bool{}

		for _, f := range strings.Split(vp.Get("LOCAL_FILESYSTEM").(string), ",") {
			locals[f] = true
		}
	}

	partitions := getPartitions()

	idx, _, err := AskForFilesystem(partitions)

	if err != nil {
		log.Fatal(err)
	}

	nFiles, err := AskForNumberOfFiles()

	if err != nil {
		log.Fatal(err)
	}

	size, err := askForSize()

	if err != nil {
		log.Fatal(err)
	}

	if confirm(nFiles, size, partitions[idx].Device, partitions[idx].MountedOn) {
		start := time.Now()
		wg := new(sync.WaitGroup)
		for i := 0; i < nFiles; i++ {
			wg.Add(1)
			fileGenerator(i, size, partitions[idx].MountedOn, wg)
		}
		wg.Wait()
		fmt.Println("Finished processing the files, the execution took", time.Since(start), "seconds")
	} else {
		fmt.Println("Aborting the command")
	}

}
