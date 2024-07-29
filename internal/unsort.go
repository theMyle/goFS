package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Unsort(root string) {
	fmt.Printf("\n--- UNSORTING ---\n\n")

	// Parse Files
	fmt.Print("PARSING FILES ")
	startTime := time.Now()

	filePaths, folderPaths, err := ScanDirRecursive(root)
	if err != nil {
		log.Fatal("error scanning directory:", err)
	}

	fmt.Printf("\t[/] ")
	fmt.Printf("\t--Time: [ %.2f ] -- Files: [ %d ] -- Folders: [ %d ] --\n",
		time.Since(startTime).Seconds(),
		len(filePaths),
		len(folderPaths))

	// Move Files
	fmt.Print("MOVING FILES ")
	startTime = time.Now()

	errMsg := "file already exists in destination:\n"
	defer func() {
		if errMsg != "" {
			fmt.Println("Some duplicated files failed to move")
		}
	}()

	for _, sourcePath := range filePaths {
		fileName := filepath.Base(sourcePath)
		destPath := filepath.Join(root, fileName)

		// Check for duplicates
		if _, err := os.Stat(destPath); err == nil {
			if sourcePath != destPath {
				errMsg += fmt.Sprintf("%v\n", sourcePath)
			}
		} else {
			if sourcePath != destPath {
				os.Rename(sourcePath, destPath)
			}
		}
	}

	fmt.Printf("\t[/] ")
	fmt.Printf("\t--Time: [ %.2f ]\n", time.Since(startTime).Seconds())

	// Clean Empty Folders
	fmt.Print("CLEAN-UP")
	startTime = time.Now()

	err = DeleteEmptyFolders(folderPaths)
	if err != nil {
		log.Fatal("error deleting empty folders:", err)
		fmt.Println("err os.rename")
	}

	fmt.Print("\t[/]")
	fmt.Printf("\t-- Time: [ %.2fs ]\n\n", time.Since(startTime).Seconds())
}
