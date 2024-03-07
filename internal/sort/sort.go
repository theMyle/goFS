package sort

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	config "github.com/theMyle/goFileSorter/internal/configs"
)

var (
	goSortFolderName = config.GoSortFolderName
	internalFolders  = config.InternalFolders
	musicsExt        = config.MusicExt
	picturesExt      = config.PicturesExt
	videosExt        = config.VideosExt
	appsExt          = config.AppsExt
	documentsExt     = config.DocumentsExt
)

func Sort(path string) {
	// Parse Files
	fmt.Println("- Parsing Files")
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	files := getFiles(entries)
	gosortFolderPath := createFolder(path, goSortFolderName)

	// Initialize map
	mp := make(map[string]string, 0)
	addMap(mp, musicsExt, internalFolders[0])
	addMap(mp, picturesExt, internalFolders[1])
	addMap(mp, videosExt, internalFolders[2])
	addMap(mp, appsExt, internalFolders[3])
	addMap(mp, documentsExt, internalFolders[4])

	// Move Files
	fmt.Println("- Moving Files")
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
			value = config.InternalFolders[5]
		}
		createFolder(gosortFolderPath, value)

		origPath := filepath.Join(path, v)
		destPath := filepath.Join(path, config.GoSortFolderName, value, v)
		fmt.Printf("\t%s\n", destPath)

		os.Rename(origPath, destPath)
	}

	fmt.Println("- Finished Sorting!")
}

func createFolder(path string, folderName string) string {
	folder := filepath.Join(path, folderName)

	err := os.MkdirAll(folder, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return folder
}

func getFiles(files []fs.DirEntry) []string {
	list := make([]string, 0)

	for _, v := range files {
		if !v.IsDir() {
			fmt.Printf("\t%s\n", v.Name())
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
