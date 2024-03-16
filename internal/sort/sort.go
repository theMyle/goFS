package sort

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	config "github.com/theMyle/goFileSorter/internal/config"
	"github.com/theMyle/goFileSorter/internal/directory"
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
	fmt.Printf("\n--- SORTING ---\n\n")
	startTime := time.Now()

	// Parse Files
	fmt.Print("PARSING FILES")
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Sort:", err)
	}

	files := GetFiles(entries)
	gosortFolderPath := directory.CreateFolder(path, goSortFolderName)
	fmt.Printf("\t[/]")
	fmt.Printf("\t-- Files: [ %v ]\n", len(files))

	// Initialize map
	mp := make(map[string]string, 0)
	addMap(mp, musicsExt, internalFolders[0])
	addMap(mp, picturesExt, internalFolders[1])
	addMap(mp, videosExt, internalFolders[2])
	addMap(mp, appsExt, internalFolders[3])
	addMap(mp, documentsExt, internalFolders[4])

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
			value = config.InternalFolders[5]
		}
		directory.CreateFolder(gosortFolderPath, value)

		origPath := filepath.Join(path, v)
		destPath := filepath.Join(path, config.GoSortFolderName, value, v)
		os.Rename(origPath, destPath)
	}
	fmt.Println("\t[/]")

	// Finish
	Finish(startTime)
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
