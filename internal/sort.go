package internal

import (
	"fmt"
	"log"
	"path/filepath"
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
	fmt.Printf("\t-- Time: [ %.2fs ] -- Files: [ %d ] -- Folders: [ %d ] --\n",
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

	for _, item := range InternalFolders {
		folder := CreateFolder(gosortFolderPath, item)
		folderPaths = append(folderPaths, folder)
	}

	// Move Files
	fmt.Print("MOVING FILES")
	startTime = time.Now()

	errMsg := ""
	defer func() {
		if errMsg != "" {
			fmt.Println("Some duplicated files failed to be moved")
		}
	}()

	for _, v := range filePaths {
		fileName := filepath.Base(v)
		extension, _ := GetFileExt(v)

		value, exists := mp[extension]
		if !exists {
			value = InternalFolders[5]
		}

		destPath := filepath.Join(gosortFolderPath, value, fileName)

		err := MoveFile(v, destPath)
		if err != nil {
			errMsg += fmt.Sprintf("%v\n", err)
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
