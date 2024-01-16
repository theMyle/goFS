package go_sort

import (
	"fmt"

	"github.com/theMyle/goFileSorter/internal/helper"
	"github.com/theMyle/goFileSorter/internal/structure"
)

func Sort() {
	var user_input string

	name, path := helper.Select_folder()

	for {
		fmt.Print("Do you wish to \"SORT\" this directory? (y or n): ")
		fmt.Scanln(&user_input)
		
		if user_input == "n" || user_input == "N" {
			return
		}
		if user_input == "y" || user_input == "Y" {
			break
		}
	}

	fmt.Printf("\nScanning Directory...\n\n")

	var root_folder structure.Folder
	structure.Init_root(&root_folder, name, path)
	structure.Scan_dir(&root_folder)
	
}