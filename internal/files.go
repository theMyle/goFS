package internal

import (
	"bufio"
	"io"
	"log"
	"os"
)

func CopyFile(src string, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		log.Fatal("error creating file: ", err)
	}
	defer destination.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(destination)

	_, err = io.Copy(writer, reader)
	if err != nil {
		log.Fatal("error copying from buffers: ", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("error flushing buffered writer: ", err)
	}

	return nil
}
