package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Sort(basePath string) {
	fmt.Printf("\n--- SORTING ---\n\n")

	// Parse Files
	fmt.Print("PARSING FILES")
	entries, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal("Sort:", err)
	}

	files := GetFiles(entries)
	gosortFolderPath := CreateFolder(basePath, GoSortFolderName)

	fmt.Printf("\t[/]")
	fmt.Printf("\t-- Files: [ %v ]\n", len(files))

	// Initialize map
	mp := make(map[string]string, 0)
	addMap(mp, MusicsExt, InternalFolders[0])
	addMap(mp, PicturesExt, InternalFolders[1])
	addMap(mp, VideosExt, InternalFolders[2])
	addMap(mp, AppsExt, InternalFolders[3])
	addMap(mp, DocumentsExt, InternalFolders[4])

	// Move Files
	fmt.Print("MOVING FILES")
	for _, v := range files {
		fileName := filepath.Ext(v)
		var extension string

		if len(fileName) > 0 {
			extension = strings.ToLower(filepath.Ext(v)[1:])
		} else {
			extension = "None"
		}

		value, exists := mp[extension]
		if !exists {
			value = InternalFolders[5]
		}
		CreateFolder(gosortFolderPath, value)

		origPath := filepath.Join(basePath, v)
		destPath := filepath.Join(basePath, GoSortFolderName, value, v)

		err := os.Rename(origPath, destPath)
		if err != nil {
			log.Fatal("error moving file: ", err)
		}
	}

	fmt.Println("\t[/]")
}

func GetFiles(files []fs.DirEntry) []string {
	list := make([]string, 0)

	for _, v := range files {
		if !v.IsDir() {
			list = append(list, v.Name())
		}
	}

	return list
}

func addMap(mp map[string]string, keys []string, value string) {
	for _, v := range keys {
		mp[v] = value
	}
}
