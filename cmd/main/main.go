package main

import (
	"github.com/theMyle/goFileSorter/internal/sort"
	"github.com/theMyle/goFileSorter/internal/unsort"
)

func main() {
	sort.Sort("C:\\Users\\jangk\\Downloads")
	unsort.Unsort("C:\\Users\\jangk\\Downloads")
}
