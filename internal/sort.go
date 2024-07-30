package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Sort(root string) {
	fmt.Printf("\n--- SORTING ---\n\n")

	// Parse Files
	fmt.Print("PARSING FILES")
	startTime := time.Now()

	filePaths, folderPaths, err := ScanDirRecursive(root)
	if err != nil {
		log.Fatal("error scanning directory:", err)
	}

	fmt.Printf("\t[/]")
	fmt.Printf("\t-- Time: [ %.2f ] -- Files: [ %d ] -- Folders: [ %d ] --\n",
		time.Since(startTime).Seconds(),
		len(filePaths),
		len(folderPaths))

	gosortFolderPath := CreateFolder(root, GoSortFolderName)

	// Initialize map
	mp := make(map[string]string, 0)
	addMap(mp, MusicsExt, InternalFolders[0])
	addMap(mp, PicturesExt, InternalFolders[1])
	addMap(mp, VideosExt, InternalFolders[2])
	addMap(mp, AppsExt, InternalFolders[3])
	addMap(mp, DocumentsExt, InternalFolders[4])

	// Move Files
	fmt.Print("MOVING FILES")
	startTime = time.Now()

	errMsg := "Files already exists in destination:\n"
	defer func() {
		if errMsg != "" {
			fmt.Println("Some duplicated files failed to move")
		}
	}()

	for _, v := range filePaths {
		fileName := filepath.Base(v)
		var extension string

		if len(filepath.Ext(fileName)) > 0 {
			extension = strings.ToLower(filepath.Ext(fileName)[1:])
		} else {
			extension = "None"
		}

		value, exists := mp[extension]
		if !exists {
			value = InternalFolders[5]
		}

		CreateFolder(gosortFolderPath, value)
		destPath := filepath.Join(root, GoSortFolderName, value, fileName)

		// check if file is already present
		if _, err := os.Stat(destPath); err == nil {
			if v != destPath {
				errMsg += fmt.Sprintf("%v\n", v)
			}
		} else {
			if v != destPath {
				os.Rename(v, destPath)
			}
		}
	}

	fmt.Print("\t[/]")
	fmt.Printf("\t-- Time: [ %.2fs ]\n", time.Since(startTime).Seconds())

	// Cleanup
	fmt.Print("CLEAN-UP")
	startTime = time.Now()

	err = DeleteEmptyFolders(folderPaths)
	if err != nil {
		log.Fatal("error deleting empty folders:", err)
	}

	fmt.Print("\t[/]")
	fmt.Printf("\t-- Time: [ %.2fs ]\n\n", time.Since(startTime).Seconds())
}

func addMap(mp map[string]string, keys []string, value string) {
	for _, v := range keys {
		mp[v] = value
	}
}
