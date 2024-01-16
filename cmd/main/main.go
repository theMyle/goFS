package main

import (
	"fmt"

	"github.com/theMyle/goFileSorter/internal/go_sort"
	"github.com/theMyle/goFileSorter/internal/helper"
	"github.com/theMyle/goFileSorter/internal/menu"
)

func main() {
	var user_input string
	var counter int = 0
	
	menu.Print_menu()

	for {
		fmt.Scanln(&user_input)

		switch user_input {
		case "1":
			helper.Clear_screen()
			go_sort.Sort()	// sort function
			// menu.Print_menu()
		case "2":
			helper.Clear_screen()
			go_sort.Un_sort() // un sort function
			// menu.Print_menu()
		case "3":
			fmt.Println("Opening Options...")
			go_sort.Options()
		case "Q", "q":
			return;
		default:
			helper.Clear_screen()
			menu.Print_menu()
			fmt.Print("\nInvalid Option. Try again. : ")
		}
		counter++
	}
}