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
	"github.com/theMyle/goFileSorter/test"
)

func main() {
	// AppRun()

	test.ScanDir("C:\\Users\\jangk\\Documents\\Games")
}

func AppRun() {
	printHeader()
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
			os.Exit(0)
		case "Q":
			os.Exit(0)
		default:
			fmt.Println("Please try again")
		}
	}
}

func printMenu() {
	fmt.Println("Select an option: (press Q to quit)")
	fmt.Println("\t1. sort")
	fmt.Println("\t2. unsort")
	fmt.Println("\t3. help")
}

func printHeader() {
	fmt.Printf("___ Go File Sorter v1.0 ___\n\n")
}

func printHelp() {
	fmt.Printf("\n-- HELP --\n\n")
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

	fmt.Printf("\nDo you really wish to (%s) this directory? (y/n)\n", strings.ToUpper(operation))
	fmt.Printf("%s: ", path)

	var input string
	fmt.Scan(&input)

	input = strings.ToLower(input)

	if input == "y" || input == "yes" {
		return path
	} else {
		os.Exit(0)
	}

	return path
}
