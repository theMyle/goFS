package go_sort

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/theMyle/goFileSorter/internal/folder"
	"github.com/theMyle/goFileSorter/internal/helper"
)

func Sort() {
	var App_Folder = "Go_Sort"
	var user_input string

	name, abs_path := helper.Select_folder()

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

	// Scan Dir
	fmt.Printf("\nScanning Directory...\n\n")
	root := folder.New_Folder(name, abs_path)
	ext := folder.Dir_Scan(&root)

	// Sorting
	fmt.Printf("\nSorting...\n\n")

	err := os.Mkdir(App_Folder, 0777)
	if err != nil {
		fmt.Printf("Using an Existing Go_Sort Folder\n\n") 
	}

	for _,v := range ext {
		folder := filepath.Join(App_Folder, v) 
		os.Mkdir(folder, 0777)

	}

	ext = []string{}
	list, _ := os.ReadDir(filepath.Join(abs_path, App_Folder))
	for _,v := range list {
		if v.IsDir() {
			ext = append(ext, v.Name())
		}
	}

	// Move Items
	folder.Move_Items(&root, abs_path)
}