package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/theMyle/goFileSorter/cmd"
)

var (
	path   string
	sort   bool
	unsort bool
	filter string
	help   bool
)

func main() {
	if len(os.Args) < 1 {
		cmd.DefaultRun()
	}

	flag.BoolVar(&sort, "sort", false, "Sorts all files in the chosen directory.")
	flag.BoolVar(&unsort, "unsort", false, "Unsorts all files in the chosen directory.")
	flag.StringVar(&filter, "filter", "copy", "Filters the chosen directory for extensions specified (supports copy or move) [default: copy]")
	flag.StringVar(&path, "path", "", "Specify the directory to be operated.")
	flag.BoolVar(&help, "h", false, "Show help message.")
	flag.BoolVar(&help, "help", false, "Show help message.")
	flag.Parse()

	if help {
		showHelp()
		os.Exit(0)
	}

	if path == "" {
		fmt.Println("error: path is required, use --help for more info.")
		os.Exit(1)
	}

	// check if path is valid
	dir, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("path error:", err)
	}
	if _, err := os.Stat(dir); err != nil {
		fmt.Println("Path supplied error: The first argument must be a path\nPlease use the -h flag for help")
	}
}

func showHelp() {
	command := make([]string, 0)
	help := make([]string, 0)

	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == "h" {
			return
		}
		command = append(command, f.Name)
		help = append(help, f.Usage)
	})
	cmd.EqualizeString(command)

	fmt.Printf("Go File Sorter by theMyle\n\n")

	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgofilesorter [FLAG] [ARGS] [ADDITIONAL ARGS..]\n\n")

	fmt.Printf("Example:\n\n")
	fmt.Printf("\tgofilesorter -path ./Downloads -sort\n")
	fmt.Printf("\tgofilesorter -filter -path ./Documents doc pdf ppt\n\n")

	fmt.Printf("Commands:\n\n")

	for i := range command {
		fmt.Printf("\t-%s\t%s\n\n", command[i], help[i])
	}
}
