package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theMyle/goFS/internal"
)

var subCmdUnsort = &cobra.Command{
	Use:     "unsort",
	Short:   programName + " - Unsorts the chosen directory",
	Long:    programName + " Unsort - Unsorts all files in the selected directory",
	Example: "  " + programName + " unsort <directory>",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error un-sorting: no directory supplied (directory is required)")
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

		var confirmation string
		fmt.Printf("Path: [%s]\nAre you sure you want to un-sort this directory? (y\\n): ", directory)
		fmt.Scanln(&confirmation)

		switch strings.ToLower(confirmation) {
		case "y":
			internal.Unsort(directory)
		case "n":
			return
		default:
			return
		}
	},
}
