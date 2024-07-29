package internal

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func Filter(path string, flagInput string, extensions []string) {
	copyFlag := false
	moveFlag := false

	// Setting the filter mode (move or copy)
	checkFlag(flagInput, "move", &moveFlag)
	checkFlag(flagInput, "copy", &copyFlag)

	fmt.Printf("\n-- FILTERING --\n\n")
	startTime := time.Now()

	// Parse Files
	fmt.Printf("PARSING FILES")
	files := parseFiles(path, extensions)
	totalFiles := 0
	for _, v := range files {
		totalFiles += len(v)
	}
	fmt.Printf("\t[/]")
	fmt.Printf("\t-- Files: [ %v ] -- Time: [ %v ]\n", totalFiles, time.Since(startTime))

	if totalFiles != 0 {
		// Move or Copy files
		var moveOrCopy func(string, string)

		if copyFlag {
			fmt.Printf("COPYING FILES")
			moveOrCopy = copyFile
		}

		if moveFlag {
			fmt.Printf("MOVING FILES")
			moveOrCopy = moveFile
		}

		goFilterFolder := CreateFolder(path, goFilter)

		var wg sync.WaitGroup
		for _, list := range files {
			wg.Add(1)
			go func() {
				extension := strings.ToLower(filepath.Ext(list[0])[1:])
				destPath := CreateFolder(goFilterFolder, extension)
				for _, file := range list {
					moveOrCopy(file, destPath)
				}
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Printf("\t[/]")
		fmt.Printf("\t-- Time: [ %v ]\n", time.Since(startTime))
	} else {
		log.Fatal("NO FILES FOUND")
	}

	// Finish
	Finish(startTime)
}

func checkFlag(input string, check string, flag *bool) {
	if strings.ToLower(input) == check {
		*flag = true
	}
}

func parseFiles(path string, extensions []string) (results [][]string) {
	files, _, _ := ScanDirRecursive(path)

	// Create a unique map
	uniqueMap := make(map[string]bool)
	for _, ext := range extensions {
		uniqueMap[ext] = true
	}

	// Create a map of string slices
	extensionsMap := make(map[string][]string)
	for _, ext := range extensions {
		extensionsMap[ext] = make([]string, 0)
	}

	// Append files with the appropriate extensions
	for _, file := range files {
		fileExt := filepath.Ext(file)
		if fileExt != "" {
			fileExt = fileExt[1:]
			if uniqueMap[fileExt] {
				extensionsMap[fileExt] = append(extensionsMap[fileExt], file)
			}
		}

	}

	// Append all to the results array
	for _, slice := range extensionsMap {
		results = append(results, slice)
	}

	return results
}

func copyFile(filePath string, dir string) {
	destPath := renameIfFileExists(filePath, dir)

	// Open source file
	source, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Filter - copyFile: Source File Error: ", err)
	}
	defer source.Close()

	// Create a destination file
	destination, err := os.Create(destPath)
	if err != nil {
		log.Fatal("Filter - copyFile: Destination File Error: ", err)
	}
	defer destination.Close()

	// Copy the contents
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal("FILTER - copyFile: io.Copy Error: ", err)
	}
}

func moveFile(filePath string, dir string) {
	destPath := renameIfFileExists(filePath, dir)

	// Move file
	err := os.Rename(filePath, destPath)
	if err != nil {
		log.Fatal("FILTER - moveFile: os.Rename Error:", err)
	}
}

func renameIfFileExists(filePath string, destPath string) (newDestPath string) {
	_, fileName := filepath.Split(filePath)
	newDestPath = filepath.Join(destPath, fileName)

	if filePath == destPath {
		return newDestPath
	}

	if _, err := os.Stat(newDestPath); err == nil {
		base := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		ext := filepath.Ext(fileName)
		counter := 1
		for {
			altName := fmt.Sprintf("%s (%d)%s", base, counter, ext)
			altPath := filepath.Join(destPath, altName)
			if _, err := os.Stat(altPath); os.IsNotExist(err) {
				newDestPath = altPath
				break
			}
			counter++
		}
	}

	return newDestPath
}
