package internal

import (
	"fmt"
	"os"
	"time"
)

func Finish(startTime time.Time) {
	elapsedTime := time.Since(startTime)
	fmt.Println("Finished \t[/]")
	fmt.Printf("\nTotal execution time: [ %v ]\n", elapsedTime)
}

func WriteLog(logPath string, message string) error {
	file, err := os.Create(logPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}
