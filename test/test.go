package test

import (
	"fmt"
	"math"
	"sync"

	"github.com/theMyle/goFileSorter/internal/directory"
	file "github.com/theMyle/goFileSorter/internal/file"
)

func ScanDir(path string) {
	files, folders := directory.ScanDir(path)

	fmt.Printf("Number of:\n")
	fmt.Printf("\tFile Paths: %d\n", len(files))
	fmt.Printf("\tDir: %d\n", len(folders))
}

func Test(files []string) {
	unique, duplicated := file.FilterDuplicated(files)
	fmt.Println("Unique:", len(unique))
	fmt.Println("Duplicated:", len(duplicated))
}

func SplitSlice(num int, list []string) (result [][]string) {
	sliceSize := int(math.Floor(float64(len(list)) / float64(num)))

	start := 0
	for i := 0; i < num; i++ {
		end := start + sliceSize
		if i == num-1 {
			result = append(result, list[start:])
			return result
		}
		result = append(result, list[start:end])
		start = end
	}

	return result
}

func Move(files []string, destPath string) {
	var wg sync.WaitGroup

	// Separate files
	unique, duplicated := file.FilterDuplicated(files)
	fmt.Printf("\t-- Unique: [ %v ] --", len(unique))
	fmt.Printf("\t-- Duplicated: [ %v ] --\n", len(duplicated))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range unique {
			file.MoveFile(v, destPath)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		file.MoveFileSlow(duplicated, destPath)
	}()

	wg.Wait()
}
