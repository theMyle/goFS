package internal

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

const bufferSize uint32 = 64 * 1024

// copy source file to destination
// ../log/sample.txt -> ../otherDir/sample.txt
func CopyFile(srcPath string, destPath string) error {
	if _, err := os.Stat(destPath); err == nil {
		return errors.New("file already exists")
	}

	source, err := os.Open(srcPath)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	defer source.Close()

	destination, err := os.Create(destPath)
	if err != nil {
		log.Fatal("error creating file: ", err)
	}
	defer destination.Close()

	buf := make([]byte, bufferSize)
	_, err = io.CopyBuffer(source, destination, buf)
	if err != nil {
		log.Fatal("error copying from buffers: ", err)
	}

	return nil
}

// move source file to destination.
// ../log/sample.txt -> ../otherDir/sample.txt
func MoveFile(srcPath string, destPath string) error {
	if _, err := os.Stat(destPath); err != nil {
		if srcPath != destPath {
			err := os.Rename(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("file already exists")
	}

	return nil
}

// returns the file extension of a file
func GetFileExt(filePath string) (string, error) {
	if len(filepath.Ext(filePath)) > 0 {
		extension := filepath.Ext(filePath)[1:]
		return extension, nil
	}
	return "", errors.New("no file extension")
}
