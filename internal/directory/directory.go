package directory

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func CleanUp(dir string) {
	entries, _ := os.ReadDir(dir)

	folders := make([]string, 0)

	for _, v := range entries {
		absPath := filepath.Join(dir, v.Name())
		if v.IsDir() {
			CleanUp(absPath)

			if IsEmpty(absPath) {
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

func IsEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(entries) == 0
}

func ScanRoot(path string) ([]string, []string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, 0)
	folders := make([]string, 0)

	for _, v := range entries {
		absPath := filepath.Join(path, v.Name())
		if v.IsDir() {
			folders = append(folders, absPath)
		}
	}

	return files, folders
}

func ScanDirRecursive(path string) (fileNameSlice []string, filesSlice []string, folderSlice []string) {
	// Create Channels
	channelBuffer := 1
	filesChan := make(chan string, channelBuffer)
	foldersChan := make(chan string, channelBuffer)
	fileNamesChan := make(chan string, channelBuffer)

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	// Recursive Goroutine
	wg.Add(1)
	go scanDirRecursive(path, filesChan, foldersChan, fileNamesChan, &wg)

	// Return Values
	files := make([]string, 0)
	folders := make([]string, 0)
	fileNames := make([]string, 0)

	// Channel Goroutines
	wg2.Add(3)
	go func() {
		defer wg2.Done()
		for file := range filesChan {
			files = append(files, file)
		}
	}()
	go func() {
		defer wg2.Done()
		for folder := range foldersChan {
			folders = append(folders, folder)
		}
	}()
	go func() {
		defer wg2.Done()
		for fileName := range fileNamesChan {
			fileNames = append(fileNames, fileName)
		}
	}()

	// Cleanup
	wg.Wait()
	close(filesChan)
	close(foldersChan)
	close(fileNamesChan)
	wg2.Wait()

	return fileNames, files, folders
}

func scanDirRecursive(path string, filesChan chan string, foldersChan chan string, fileNamesChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	subDir := make([]string, 0)

	// Take the filepath, folderpath, and filenames
	for _, item := range entries {
		absPath := filepath.Join(path, item.Name())
		if item.IsDir() {
			subDir = append(subDir, absPath)
			foldersChan <- absPath
		} else {
			filesChan <- absPath
			fileNamesChan <- item.Name()
		}
	}

	// Recurse
	for _, dir := range subDir {
		wg.Add(1)
		go scanDirRecursive(dir, filesChan, foldersChan, fileNamesChan, wg)
	}
}
