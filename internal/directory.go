package internal

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
)

// Check if dir is empty
func IsEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("error reading directory: ", err)
	}
	return len(entries) == 0
}

// Scans root dir recursively and returns a list of file and folder paths.
func ScanDirRecursive(root string) (files []string, folders []string, Error error) {
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			folders = append(folders, path)
		} else {
			files = append(files, path)
		}
		return nil

	})
	if err != nil {
		return []string{}, []string{}, err
	}

	return files, folders, nil
}

// checks and delete empty folders recursively
func DeleteEmptyFolders(folders []string) error {
	slices.Sort(folders)
	slices.Reverse(folders)

	for _, v := range folders {
		dir, err := os.ReadDir(v)
		if err != nil {
			return err
		}

		if len(dir) == 0 {
			err := os.Remove(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateFolder(path string, folderName string) string {
	folder := filepath.Join(path, folderName)

	err := os.MkdirAll(folder, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return folder
}
