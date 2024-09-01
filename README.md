# Go File Sorter

**goFS** - A fast concurrent file sorter

## Features

- SORT: Sorts all files in the chosen directory.
- UNSORT: Unsorts all files in the chosen directory.
- FILTER: Filters files with the specified file extensions inside chosen directory.

# CLI 

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

# Installation

If you have go compiler installed 

```bash
go install github.com/theMyle/goFS@latest
```
You can now launch the app by simply typing `goFS` in the commandline.

Or just download the precompiled binary.

# Notes

- if a file with same name already exists in the destination path, the file will not be moved and will therfore stay there. 
