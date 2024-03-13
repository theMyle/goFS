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
)

func main() {
	AppRun()
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
			sort.Unsort(path)
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
	fmt.Printf("<____> Go File Sorter v1.0 <____>\n")
}

func printHelp() {
	fmt.Printf("\n-- HELP --\n")
	fmt.Println("\tsort - ")
	fmt.Printf("\tunsort - \n\n")
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
