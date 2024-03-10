package unsort

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/theMyle/goFileSorter/internal/directory"
	fle "github.com/theMyle/goFileSorter/internal/files"
)

func Unsort(root string) {
	fmt.Printf("\n--- UNSORTING ---\n\n")

	startTime := time.Now()

	// Look for files
	fmt.Print("Parsing Files ")
	files, folders := directory.Scan(root)

	for _, dir := range folders {
		directory.ScanRecursive(dir, &files, &folders)
	}

	fmt.Printf("\t[/] ")
	fmt.Printf("\t- %d files\n", len(files))

	// move files
	fmt.Print("Moving Files ")
	fle.MoveFiles(files, root)
	fmt.Println("\t[/]")

	// clean empty folder
	fmt.Print("Cleanup ")
	for _, v := range folders {
		cleanUp(v)
	}
	fmt.Println("\t[/]")

	elapsedTime := time.Since(startTime)
	fmt.Println("Finished \t[/]")
	fmt.Printf("\nTotal execution time: (%v)\n", elapsedTime)
}

func cleanUp(dir string) {
	entries, _ := os.ReadDir(dir)

	folders := make([]string, 0)

	for _, v := range entries {
		absPath := filepath.Join(dir, v.Name())
		if v.IsDir() {
			cleanUp(absPath)

			if directory.IsEmpty(absPath) {
				folders = append(folders, absPath)
			}
		}
	}

	for _, folder := range folders {
		os.Remove(folder)
	}

	entries, _ = os.ReadDir(dir)
	if len(entries) == 0 {
		os.Remove(dir)
	}
}
