package main

import (
	"fmt"
	"os"
)

func main() {

	dir, _ := os.Getwd()
	entries, _ := os.ReadDir(dir)

	fmt.Println(dir)
	fmt.Println("List of items in this directory: ", entries)

	for _, entry := range entries {
		fmt.Println(entry.Name())
		fmt.Println(entry.IsDir())
	}
}