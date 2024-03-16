package sort

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/theMyle/goFileSorter/internal/directory"
)

func Unsort(root string) {
	fmt.Printf("\n--- UNSORTING ---\n\n")
	startTime := time.Now()

	// Parse Files
	fmt.Print("PARSING FILES ")

	files, folders := directory.ScanDir(root)

	fmt.Printf("\t[/] ")
	fmt.Printf("\t-- Files: [ %d ] -- Folders: [ %d ] -- Time: [ %v ] --\n", len(files), len(folders), time.Since(startTime))

	// Move Files
	fmt.Print("MOVING FILES ")
	move(files, root)
	fmt.Print("\t[/] ")
	fmt.Printf("\t-- Time: [ %v ]\n", time.Since(startTime))

	// Clean Empty Folders
	fmt.Print("CLEANUP ")
	for _, v := range folders {
		directory.CleanUp(v)
	}
	fmt.Println("\t[/]")

	// Finish
	Finish(startTime)
}

func move(files []string, destPath string) {
	var wg sync.WaitGroup

	// Separate files
	unique, duplicated := filterDuplicated(files)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, absPath := range unique {
			moveFileBasic(absPath, destPath)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		moveFileComplex(duplicated, destPath)
	}()

	wg.Wait()
}

func moveFileComplex(files []string, destPath string) {
	for _, absPath := range files {
		file, err := os.Stat(absPath)
		if err != nil {
			log.Fatal(err)
		}

		name := file.Name()
		newPath := filepath.Join(destPath, name)

		if absPath == newPath {
			return
		}

		if _, err := os.Stat(newPath); err == nil {
			base := strings.TrimSuffix(name, filepath.Ext(name))
			ext := filepath.Ext(name)
			counter := 1
			for {
				altName := fmt.Sprintf("%s (%d)%s", base, counter, ext)
				altPath := filepath.Join(destPath, altName)
				if _, err := os.Stat(altPath); os.IsNotExist(err) {
					newPath = altPath
					break
				}
				counter++
			}
		}

		err = os.Rename(absPath, newPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func moveFileBasic(absPath string, destPath string) {
	_, name := filepath.Split(absPath)
	newPath := filepath.Join(destPath, name)

	if absPath == newPath {
		return
	}

	err := os.Rename(absPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func filterDuplicated(files []string) (unique []string, hasCopy []string) {
	uniqueMap := make(map[string]bool)

	for _, file := range files {
		_, fileName := filepath.Split(file)
		if uniqueMap[fileName] {
			uniqueMap[fileName] = false
		} else {
			uniqueMap[fileName] = true
		}
	}

	for _, file := range files {
		_, fileName := filepath.Split(file)
		if uniqueMap[fileName] {
			unique = append(unique, file)
		} else {

			hasCopy = append(hasCopy, file)
		}
	}

	return unique, hasCopy
}
