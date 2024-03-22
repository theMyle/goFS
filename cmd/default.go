package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/theMyle/goFileSorter/internal/sort"
)

func DefaultRun() {
	printHeader()
	printMenu()

	var userInput string
	var loop bool = true
	reader := bufio.NewReader(os.Stdin)

	for loop {
		fmt.Printf(">> ")
		fmt.Scan(&userInput)
		switch userInput {
		case "1":
			path := getFolder("sort")
			sort.Sort(path)
			loop = false
			reader.ReadString('\n')
		case "2":
			path := getFolder("unsort")
			sort.Unsort(path)
			loop = false
			reader.ReadString('\n')
		case "3":
			filterFlag := "copy"
			path := getFolder("filter & copy")
			extensions := getFileExtensions()
			sort.Filter(path, filterFlag, extensions)
			loop = false
		case "4":
			filterFlag := "move"
			path := getFolder("filter & move")
			extensions := getFileExtensions()
			sort.Filter(path, filterFlag, extensions)
			loop = false
		case "5":
			printHelp()
			printMenu()
		case "q":
			os.Exit(0)
		case "Q":
			os.Exit(0)
		default:
			fmt.Print("Please try again")
		}
	}

	fmt.Println("Press Enter to exit...")
	reader.ReadString('\n')
}

func printMenu() {
	fmt.Println("Select an option: (press Q to quit)")
	fmt.Println("\t1. sort")
	fmt.Println("\t2. unsort")
	fmt.Println("\t3. filter & copy")
	fmt.Println("\t4. filter & move")
	fmt.Println("\t5. help")
}

func printHeader() {
	fmt.Printf("<____> Go File Sorter v1.0 <____>\n\n")
}

func printHelp() {
	fmt.Printf("\n-- HELP --\n")

	fmt.Println("1. [ sort ]: ")
	fmt.Printf("\t-- sorts the files inside the chosen directory (not including ones inside folders).\n\n")

	fmt.Println("2. [ unsort ]: ")
	fmt.Printf("\t-- unsorts all files and folders inside the chosen directory.\n\n")

	fmt.Println("3. [ filter & copy ]: ")
	fmt.Printf("\t-- filters all files with the specified extension and creates a copy in a separate directory.\n\n")

	fmt.Println("4. [ filter & move ]: ")
	fmt.Printf("\t-- filters all files with the specified extension and moves it into a separated directory.\n\n")
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

func getFileExtensions() []string {
	fmt.Println("\nEnter the file extension\\s the you wish to filter: ")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	ext, _ := reader.ReadString('\n')

	extensions := strings.Split(strings.TrimSpace(ext), " ")

	for i, ext := range extensions {
		extensions[i] = strings.ToLower(ext)
	}

	return extensions
}
