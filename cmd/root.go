package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	programName string = "goFS"
)

func init() {
	rootCmd.AddCommand(subCmdSort)
	rootCmd.AddCommand(subCmdUnsort)
	rootCmd.AddCommand(subCmdFilter)
}

var rootCmd = &cobra.Command{
	Use:   programName + " [sort | unsort | filter] <directory> [flags] [args]",
	Short: "Go File Sorter by theMyle",
	Long:  programName + " - A fast concurrent file sorter made with golang",
	Example: "  " + programName + " sort ./Downloads\n  " +
		programName + " unsort ./Documents\n  " +
		programName + " filter --copy ./Documents exe pdf doc",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
