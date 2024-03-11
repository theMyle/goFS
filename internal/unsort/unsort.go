package unsort

import (
	"fmt"
	"time"

	"github.com/theMyle/goFileSorter/internal/directory"
	fle "github.com/theMyle/goFileSorter/internal/files"
)

func Unsort(root string) {
	fmt.Printf("\n--- UNSORTING ---\n\n")
	startTime := time.Now()

	// Parse Files
	fmt.Print("Parsing Files ")
	_, files, folders := directory.ScanDirRecursive(root)
	fmt.Printf("\t[/] ")
	fmt.Printf("\t- %d files\n", len(files))

	// Move Files
	fmt.Print("Moving Files ")
	fle.MoveFiles(files, root)
	fmt.Println("\t[/]")

	// Clean Empty Folders
	fmt.Print("Cleanup ")
	for _, v := range folders {
		directory.CleanUp(v)
	}
	fmt.Println("\t[/]")

	// Finish
	elapsedTime := time.Since(startTime)
	fmt.Println("Finished \t[/]")
	fmt.Printf("\nTotal execution time: (%v)\n", elapsedTime)
}
