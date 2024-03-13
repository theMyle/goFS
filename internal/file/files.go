package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FilterDuplicated(files []string) (unique []string, hasCopy []string) {
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

func MoveFileSlow(files []string, destPath string) {
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

func MoveFile(absPath string, destPath string) {
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

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
