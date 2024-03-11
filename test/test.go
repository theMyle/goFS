package test

import (
	"fmt"
	"time"

	"github.com/theMyle/goFileSorter/internal"
	"github.com/theMyle/goFileSorter/internal/directory"
)

func ScanDir(path string) {
	fileNames, files, folders := directory.ScanDirRecursive(path)

	fmt.Printf("Number of:\n")
	fmt.Printf("\tFiles: %d\n", len(fileNames))
	fmt.Printf("\tFile Paths: %d\n", len(files))
	fmt.Printf("\tDir: %d\n", len(folders))
}

func MakeUnique(files []string) {
	start := time.Now()
	internal.MakeUniqueMap(files)
	elapsedTime := time.Since(start)
	fmt.Printf("Exec Time: %v\n", elapsedTime)

	start = time.Now()
	internal.FilterDuplicated(internal.MakeUniqueMap(files))
	elapsedTime = time.Since(start)
	fmt.Printf("Exec Time: %v\n", elapsedTime)
}


