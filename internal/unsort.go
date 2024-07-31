package internal

import (
	"fmt"
	"log"
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
	fmt.Printf("\t-- Time: [ %.2fs ] -- Files: [ %d ] -- Folders: [ %d ] --\n",
		time.Since(startTime).Seconds(),
		len(filePaths),
		len(folderPaths))

	// Move Files
	fmt.Print("MOVING FILES ")
	startTime = time.Now()

	errMsg := ""
	defer func() {
		if errMsg != "" {
			fmt.Println("Some duplicated files failed to be moved")
		}
	}()

	for _, sourcePath := range filePaths {
		fileName := filepath.Base(sourcePath)
		destPath := filepath.Join(root, fileName)

		err := MoveFile(sourcePath, destPath)
		if err != nil {
			errMsg += fmt.Sprintf("%v\n", err)
		}
	}

	fmt.Printf("\t[/] ")
	fmt.Printf("\t-- Time: [ %.2fs ]\n", time.Since(startTime).Seconds())

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
