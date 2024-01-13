package main

import (
	"fmt"

	"github.com/theMyle/goFileSorter/internal/go_sort"
	"github.com/theMyle/goFileSorter/internal/helper"
)

// Project data
var version string = "v1.0"
// var date string = "01/13/2024";


// Menu Options
var option_1 string = "Sort"
var option_2 string = "Un_sort"
var option_3 string = "Settings"

func print_menu() {
	helper.Clear_screen()

	fmt.Printf("|_-_-_|     Go File Sorter %s     |_-_-_|\n\n", version)
	fmt.Printf("Choose an option down below (press Q to quit):\n")
	fmt.Printf("\t1. %s\n", option_1)
	fmt.Printf("\t2. %s\n", option_2)
	fmt.Printf("\t3. %s\n", option_3)
}

func main() {
	var user_input string
	var counter int = 0
	
	print_menu()

	for {
		fmt.Scanln(&user_input)

		switch user_input {
		case "1":
			helper.Clear_screen()
			go_sort.Sort()	// sort function
			print_menu()
		case "2":
			helper.Clear_screen()
			go_sort.Un_sort() // un sort function
			print_menu()
		case "3":
			fmt.Println("Opening Options...")
			go_sort.Options()
		case "Q", "q":
			return;
		default:
			helper.Clear_screen()
			print_menu()
			fmt.Print("\nInvalid Option. Try again. : ")
		}
		counter++
	}
}