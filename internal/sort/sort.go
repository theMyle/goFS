package sort

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var InternalFolders []string = []string{"Music", "Pictures", "Videos", "Apps", "Documents", "extras"}
var musicExt []string = []string{"mp3", "wav", "aiff", "flac", "aac", "mod", "ogg", "wma", "m4a", "amr", "mid", "midi", "ac3", "alac", "ape", "dsd", "opus", "au", "pcm", "mp4", "3gp", "ra", "rm", "vorbis", "wv", "tta", "dts", "g723", "g726", "xwm", "caf", "eac3", "mka", "mlp", "s3m", "it"}
var picturesExt []string = []string{"jpeg", "jpg", "png", "gif", "bmp", "tiff", "tif", "svg", "webp", "heif", "heic", "ico", "psd", "raw", "pcx", "pbm", "pgm", "ppm", "pnm", "webp", "jfif", "jpe", "jps", "j2k", "jp2", "jxr", "wdp", "xbm", "xpm", "eps", "pdf", "emf", "wmf", "cgm", "dcm", "dicom", "exif"}
var videosExt []string = []string{"mp4", "avi", "mkv", "mov", "wmv", "flv", "mpeg", "webm", "3gp", "m4v", "asf", "rm", "swf", "vob", "ts", "ogv", "mpg", "m2ts", "mxf", "hevc", "mts", "avchd", "flv", "mov", "wmv", "rmvb"}
var appsExt []string = []string{"exe", "apk", "app", "msi", "dmg", "deb", "rpm", "jar", "sh", "bat", "vbs", "ps1", "py", "bin", "elf"}
var documentsExt []string = []string{"docx", "doc", "pdf", "xlsx", "xls", "pptx", "ppt", "odt", "ods", "odp", "rtf", "txt", "csv", "html", "xml", "json", "md", "epub", "tex", "wpd"}

func Sort(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// Parse Files
	files := getFiles(entries)
	gosortPath := createFolder(path, GoSortFolderName)

	// Initialize Maps
	mp := make(map[string]string, 0)
	addMap(mp, musicExt, InternalFolders[0])
	addMap(mp, picturesExt, InternalFolders[1])
	addMap(mp, videosExt, InternalFolders[2])
	addMap(mp, appsExt, InternalFolders[3])
	addMap(mp, documentsExt, InternalFolders[4])

	// Move Files
	for _, v := range files {
		extension := strings.ToLower(filepath.Ext(v)[1:])

		value, exists := mp[extension]
		if !exists {
			value = InternalFolders[5]
		}
		createFolder(gosortPath, value)

		origPath := filepath.Join(path, v)
		destPath := filepath.Join(path, GoSortFolderName, value, v)

		os.Rename(origPath, destPath)
	}
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
