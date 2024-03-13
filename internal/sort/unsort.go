package sort

import (
	"fmt"
	"sync"
	"time"

	"github.com/theMyle/goFileSorter/internal/directory"
	"github.com/theMyle/goFileSorter/internal/file"
)

func Unsort(root string) {
	fmt.Printf("\n--- UNSORTING ---\n\n")
	startTime := time.Now()

	// Parse Files
	fmt.Print("PARSING FILES ")

	files, folders := directory.ScanDir(root)

	fmt.Printf("\t[/] ")
	fmt.Printf("\t-- Files: [ %d ] -- Folders: [ %d ] -- Time: [ %v ] --\n", len(files), len(folders), time.Since(startTime))

	// Move Files
	fmt.Print("MOVING FILES ")
	move(files, root)
	fmt.Print("\t[/] ")
	fmt.Printf("\t-- Time: [ %v ]\n", time.Since(startTime))

	// Clean Empty Folders
	fmt.Print("CLEANUP ")
	for _, v := range folders {
		directory.CleanUp(v)
	}
	fmt.Println("\t[/]")

	// Finish
	elapsedTime := time.Since(startTime)
	fmt.Println("Finished \t[/]")
	fmt.Printf("\nTotal execution time: [ %v ]\n", elapsedTime)
}

func move(files []string, destPath string) {
	var wg sync.WaitGroup

	// Separate files
	unique, duplicated := file.FilterDuplicated(files)

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
