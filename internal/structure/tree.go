package structure

import (
	"fmt"
	"os"
)

type Folder struct {
	FolderName string
	FolderPath string
	Files      []string
	Subfolders []*Folder
}

func Init_root(ptr *Folder, name string, path string) {
	ptr.FolderName = name
	ptr.FolderPath = path
	ptr.Files = []string{}
	ptr.Subfolders = []*Folder{}
}

func Scan_dir(ptr *Folder) {
	// var folder_name string = ptr.FolderName
	var folder_path string = ptr.FolderPath

	entries, err := os.ReadDir(folder_path)
	if err != nil {
		fmt.Println("Error Scanning Directories")
		return
	}

	for _, value := range entries {
		if value.IsDir() {
			fmt.Println("Folder: ",value.Name())
		} else {
			fmt.Println("File: ", value.Name())
		}
	}
}

