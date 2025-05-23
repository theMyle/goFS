# ⚡️ goFS – Go File Sorter

**goFS** (Go File Sorter) is a fast, concurrent CLI tool written in Go for sorting, unsorting, and filtering files in large directories. Built as a learning project, it demonstrates how concurrency can accelerate I/O-heavy tasks — though, as the project revealed, most file operations are I/O-bound, meaning performance is ultimately limited by disk read/write speeds rather than CPU power.

Through this project, I deepened my understanding of:

Go's concurrency model (goroutines, channels, sync)

How caching impacts file system performance

Real-world trade-offs between CPU-bound vs I/O-bound task

---

## 📦 Features

- **Sort** files into subfolders based on file type
- **Unsort** files by restoring them to their original (flattened) structure
- **Filter** files by extension, with the option to move or copy them
- Built with **Golang** and the **Cobra** CLI library
- Concurrent file operations for fast performance

---

## CLI 

```
goFS - A fast concurrent file sorter made with golang

Usage:
  goFS [command] [flag] [directory] [args]
  goFS [command]

Examples:
  goFS sort ./Downloads
  goFS unsort ./Documents
  goFS filter --copy ./Documents exe pdf doc

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  filter      goFS - Filters the chosen directory
  help        Help about any command
  sort        goFS - Sorts the chosen directory
  unsort      goFS - Unsorts the chosen directory

Flags:
  -h, --help   help for goFS

Use "goFS [command] --help" for more information about a command.
```

## Sample Usage

```
--- SORTING ---

PARSING FILES   [/]     -- Time: [ 2.09s ] -- Files: [ 19483 ] -- Folders: [ 3495 ] --
MOVING FILES    [/]     -- Time: [ 14.10s ]
CLEAN-UP        [/]     -- Time: [ 0.83s ]

Press Enter to exit...
_______________________________________________________________

--- UNSORTING ---

PARSING FILES   [/]     -- Time: [ 2.88s ] -- Files: [ 19483 ] -- Folders: [ 3532 ] --
MOVING FILES    [/]     -- Time: [ 14.89s ]
CLEAN-UP        [/]     -- Time: [ 0.61s ]

Press Enter to exit...
_______________________________________________________________

-- FILTERING --

PARSING FILES   [/]     -- Time: [ 3.28s ] -- Files: [ 26606 ]
COPYING FILES   [/]     -- Time: [ 3.94s ]
CLEAN-UP        [/]     -- Time: [ 0.46s ]

Press Enter to exit...
_______________________________________________________________

-- FILTERING --

PARSING FILES   [/]     -- Time: [ 2.82s ] -- Files: [ 17848 ]
MOVING FILES    [/]     -- Time: [ 7.48s ]
CLEAN-UP        [/]     -- Time: [ 0.43s ]

Press Enter to exit...
_______________________________________________________________
```

## Installation

If you have go compiler installed 

```bash
go install github.com/theMyle/goFS@latest
```

You can now launch the app by simply typing `goFS` in the commandline.

Or just download the precompiled binary.

## Notes

- if a file with same name already exists in the destination path, the file will not be moved and will therfore stay there. 
