package helper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

func Select_folder() (string, string) {
	dir, err := dialog.Directory().Title("Choose a folder to sort").Browse()
	if err != nil {
		println(err, ": Error selecting folder")
		return "", ""
	}

	substrings := strings.Split(dir, "\\")
	folder_name := substrings[len(substrings)-1]

	err = os.Chdir(dir)
	if 	(err != nil) {
		println(err, ": Error changing directory")
		return "", ""
	}

	fmt.Println(dir)

	return folder_name, dir
}

func Clear_screen() {
	if (runtime.GOOS == "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	if (runtime.GOOS == "linux") {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func User_input(prompt string) string {
	var returnVal string

	fmt.Print(prompt)
	fmt.Scanln(&returnVal)

	return returnVal
}