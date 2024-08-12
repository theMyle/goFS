package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theMyle/goFS/internal"
)

var subCmdSort = &cobra.Command{
	Use:     "sort",
	Short:   programName + " - Sorts the chosen directory",
	Long:    programName + " Sort - Sorts all files in the chosen directory",
	Example: "  " + programName + " sort ./Downloads",

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

		var confirmation string
		fmt.Printf("Path: [%s]\nAre you sure you want to sort this directory? (y\\n): ", directory)
		fmt.Scanln(&confirmation)

		switch strings.ToLower(confirmation) {
		case "y":
			internal.Sort(directory)
		case "n":
			return
		default:
			return
		}
	},
}
