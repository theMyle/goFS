package unsort

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func Unsort(root string) {
	files, folders := scanDir(root)
	var wg sync.WaitGroup

	// Scanning Files
	fmt.Println("- Scanning Directory")
	for _, v := range folders {
		wg.Add(1)
		scanDirRecursive(v, &files, &folders, &wg)
	}

	wg.Wait()

	// Moving Files
	fmt.Println("- Moving Files")
	moveFiles(files, root)

	// Cleanup
	fmt.Println("- Cleaning up folders")
	for _, v := range folders {
		fmt.Println("\t", v)
		cleanUp(v)
	}

	fmt.Println("- Finished Unsorting")
}

func scanDir(path string) ([]string, []string) {
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

func scanDirRecursive(path string, fileContainer *[]string, folderContainer *[]string, wg *sync.WaitGroup) {
	defer wg.Done()

	var subWG sync.WaitGroup

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range entries {
		absPath := filepath.Join(path, v.Name())
		fmt.Printf("\t%s\n", absPath)
		if v.IsDir() {
			*folderContainer = append(*folderContainer, absPath)
			subWG.Add(1)
			go func(dir string) {
				scanDirRecursive(dir, fileContainer, folderContainer, &subWG)
			}(absPath)
		} else {
			*fileContainer = append(*fileContainer, absPath)
		}
	}

	subWG.Wait()
}

func moveFiles(files []string, destPath string) {
	for _, v := range files {
		file, err := os.Stat(v)
		if err != nil {
			log.Fatal(err)
		}

		name := file.Name()

		currentPath := v
		newPath := filepath.Join(destPath, name)

		present := fileExists(newPath)
		altPath := newPath
		counter := 1
		for present {
			altPath = fmt.Sprintf("%s (%d)", newPath, counter)
			present = fileExists(altPath)
			counter++
		}

		newPath = altPath
		fmt.Printf("\t%s\n", newPath)

		err = os.Rename(currentPath, newPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}

func cleanUp(dir string) {
	entries, _ := os.ReadDir(dir)

	folders := make([]string, 0)

	for _, v := range entries {
		absPath := filepath.Join(dir, v.Name())
		if v.IsDir() {
			cleanUp(absPath)

			if isEmptyDir(absPath) {
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

func isEmptyDir(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(entries) == 0
}
