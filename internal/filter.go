package internal

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func Filter(path string, flagInput string, extensions []string) {
	copyFlag := false
	moveFlag := false
	extMap := stringToMap(extensions)

	// Setting the filter mode (move or copy)
	checkFlag(flagInput, "move", &moveFlag)
	checkFlag(flagInput, "copy", &copyFlag)

	fmt.Printf("\n-- FILTERING --\n\n")

	// parse files
	fmt.Printf("PARSING FILES")
	startTime := time.Now()

	filePaths, folderPaths, err := ScanDirRecursive(path)
	if err != nil {
		log.Fatal("error parsing files:", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	filteredFiles := []string{}

	for _, file := range filePaths {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if len(filepath.Ext(file)) > 0 {
				ext := strings.ToLower(filepath.Ext(file)[1:])
				if _, exists := extMap[ext]; exists {
					mu.Lock()
					filteredFiles = append(filteredFiles, file)
					mu.Unlock()
				}
			}
		}(file)
	}

	wg.Wait()
	fmt.Printf("\t[/]")
	fmt.Printf("\t-- Time: [ %.2fs ] -- Files: [ %v ]\n", time.Since(startTime).Seconds(), len(filteredFiles))

	// move or copy files
	goFilterFolderPath := CreateFolder(path, goFilter)
	for _, item := range extensions {
		CreateFolder(goFilterFolderPath, item)
	}

	var moveOrCopyFunc func(srcPath string, destPath string) error
	if copyFlag {
		fmt.Printf("COPYING FILES")
		moveOrCopyFunc = CopyFile
	} else {
		fmt.Printf("MOVING FILES")
		moveOrCopyFunc = MoveFile
	}

	startTime = time.Now()
	errMsg := ""
	defer func() {
		if errMsg != "" {
			fmt.Println("Some duplicated files failed to be copy/moved")
		}
	}()

	for _, file := range filteredFiles {
		fileName := filepath.Base(file)
		fileExt, _ := GetFileExt(file)
		destPath := filepath.Join(goFilterFolderPath, strings.ToLower(fileExt), fileName)
		err := moveOrCopyFunc(file, destPath)
		if err != nil {
			errMsg += fmt.Sprintf("%v\n", err)
		}
	}

	fmt.Printf("\t[/]")
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

func checkFlag(input string, check string, flag *bool) {
	if strings.ToLower(input) == check {
		*flag = true
	}
}

// coverts string slice into a map
func stringToMap(stringSlice []string) map[string]bool {
	mp := make(map[string]bool)

	for _, item := range stringSlice {
		if !mp[item] {
			mp[item] = true
		}
	}

	return mp
}
