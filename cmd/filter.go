package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theMyle/goFS/internal"
)

var (
	copy bool
	move bool
)

func init() {
	subCmdFilter.Flags().BoolVarP(&copy, "copy", "c", false, "copies all the filtered items")
	subCmdFilter.Flags().BoolVarP(&move, "move", "m", false, "moves all the filtered items")
}

var subCmdFilter = &cobra.Command{
	Use:     "filter [flags] <directory> [extensions]",
	Short:   programName + " - Filters the chosen directory",
	Long:    programName + " Filter - Filters all files with extensions specified inside the target directory",
	Example: "  " + programName + " filter --copy ./Documents exe pdf docs",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		directory := args[0]
		directory, err := filepath.Abs(directory)
		if err != nil {
			fmt.Printf("Error converting to absolute path:\n\t%s\n", err)
			return
		}

		_, err = os.Stat(directory)
		if err != nil {
			fmt.Printf("Invalid path error:\n\t%s\n", err)
			return
		}

		if (copy && move) || (!copy && !move) {
			fmt.Println("Invalid filter usage: Must select one (--copy or --move)")
			return
		}

		var filterMode string
		if copy {
			filterMode = "copy"
		} else {
			filterMode = "move"
		}

		fileExtensions := args[1:]
		var confirmation string
		fmt.Printf("Path: \t\t[%s]\nExtensions: \t%v\nMode: \t\t[%s]\n\n", directory, fileExtensions, filterMode)
		fmt.Print("Are you sure you want to filter these items from this directory? (y\\n): ")
		fmt.Scanln(&confirmation)

		switch strings.ToLower(confirmation) {
		case "y":
			fmt.Println("filtering ", fileExtensions)
			internal.Filter(directory, filterMode, fileExtensions)
		default:
			return
		}
	},
}
