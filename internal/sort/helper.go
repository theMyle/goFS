package sort

import (
	"fmt"
	"time"
)

func Finish(startTime time.Time) {
	elapsedTime := time.Since(startTime)
	fmt.Println("Finished \t[/]")
	fmt.Printf("\nTotal execution time: [ %v ]\n", elapsedTime)
}
