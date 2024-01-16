package menu

import (
	"fmt"

	"github.com/theMyle/goFileSorter/internal/helper"
)

var version string = "v1.0"
// var date string = "01/13/2024";

var option_1 string = "Sort"
var option_2 string = "Un_sort"
var option_3 string = "Settings"

func Print_menu() {
	helper.Clear_screen()

	fmt.Printf("|_-_-_|     Go File Sorter %s     |_-_-_|\n\n", version)
	fmt.Printf("Choose an option down below (press Q to quit):\n")
	fmt.Printf("\t1. %s\n", option_1)
	fmt.Printf("\t2. %s\n", option_2)
	fmt.Printf("\t3. %s\n", option_3)
}