package directory

import (
	"log"
	"os"
	"path/filepath"
)

func IsEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(entries) == 0
}

func Scan(path string) ([]string, []string) {
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

func ScanRecursive(path string, fileContainer *[]string, folderContainer *[]string) {
	walkRecurse(path, fileContainer, folderContainer)
}

func walkRecurse(path string, fileContainer *[]string, folderContainer *[]string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var subDir []string

	for _, item := range entries {
		absPath := filepath.Join(path, item.Name())
		if item.IsDir() {
			subDir = append(subDir, absPath)
			*folderContainer = append(*folderContainer, absPath)
		} else {
			*fileContainer = append(*fileContainer, absPath)
		}
	}

	for _, dir := range subDir {
		walkRecurse(dir, fileContainer, folderContainer)
	}
}
