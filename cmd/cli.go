package cmd

import (
	"fmt"
	"strings"
)

func EqualizeString(StrSlice []string) {
	longest := 0

	for _, val := range StrSlice {
		len := len(val)
		if len > longest {
			longest = len
		}
	}

	for i := range StrSlice {
		StrSlice[i] = fmt.Sprintf("%s%s ", StrSlice[i], strings.Repeat(" ", longest-len(StrSlice[i])))
	}
}
