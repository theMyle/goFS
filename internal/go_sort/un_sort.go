package go_sort

import (
	"fmt"

	"github.com/theMyle/goFileSorter/internal/helper"
)

func Un_sort() {
	var user_input string

	helper.Select_folder()

	for {
		fmt.Print("Do you wish to \"UN_SORT\" this directory? (y or n): ")
		fmt.Scanln(&user_input)

		if user_input == "n" || user_input == "N" {
			return
		}
		if user_input == "y" || user_input == "Y" {
			break
		}
	}
	// un sorting algorithm
}