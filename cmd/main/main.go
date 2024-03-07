package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/theMyle/goFileSorter/internal/sort"
	"github.com/theMyle/goFileSorter/internal/unsort"
)

func main() {
	clearScreen()
	printMenu()

	var userInput string
	var loop bool = true

	for loop {
		fmt.Printf(">> ")
		fmt.Scan(&userInput)
		switch userInput {
		case "1":
			path := getFolder("sort")
			sort.Sort(path)
			loop = false
		case "2":
			path := getFolder("unsort")
			unsort.Unsort(path)
			loop = false
		case "3":
			printHelp()
			printMenu()
		case "q":
			loop = false
		default:
			fmt.Println("Please try again")
		}
	}

	fmt.Println("\nGo Sort Complete!")
	fmt.Printf("Enter (q) to Quit: ")
	fmt.Scan(&userInput)
}

func printMenu() {
	fmt.Println("Please select an option down below: (press (q) to quit)")
	fmt.Println("\t1. sort")
	fmt.Println("\t2. unsort")
	fmt.Println("\t3. help")
}

func printHeader() {
	fmt.Printf("#_-_-_# Go File Sorter v1.0 #_-_-_#\n\n")
}

func printHelp() {
	fmt.Printf("\n-- HELP --\n")
	fmt.Println("(1) - SORT: sorts the chosen directory non recursively (not including other folders). It will only sort top level files.")
	fmt.Printf("(2) - UNSORT: unsorts the entire directory recursively and move files outside to the root directory. (USE CAREFULLY)\n\n")
}

func clearScreen() {
	osName := runtime.GOOS
	var cmd *exec.Cmd

	if osName == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if osName == "linux" {
		cmd = exec.Command("clear")
	} else {
		panic("Unsupported operating system")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
	printHeader()
}

func getFolder(operation string) string {
	path, err := dialog.Directory().Title("Choose directory").Browse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nDo you really wish to %s this directory (y/n)?\n", strings.ToUpper(operation))
	fmt.Printf("%s: ", path)

	var input string
	fmt.Scan(&input)

	if input == "y" || input == "Y" {
		return path
	} else {
		os.Exit(0)
	}

	return path
}
