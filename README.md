# Go File Sorter

A simple CLI app for sorting files

## Features

- `SORT`: Sorts the entire content of the selected directory.
- `UNSORT`: Unsorts all files in the chosen directory (Puts them outside).
- `FILTER`: Filters files with the specified file extensions in the chosen directory.
  - `move` - moves files.
  - `copy` - copies files.

# Usage

```
Go File Sorter by theMyle

Select an option: (enter Q to quit)
        1. sort
        2. unsort
        3. filter & copy
        4. filter & move
        5. help
>>
```

## Sort

```
--- SORTING ---

PARSING FILES   [/]     -- Time: [ 2.09s ] -- Files: [ 19483 ] -- Folders: [ 3495 ] --
MOVING FILES    [/]     -- Time: [ 14.10s ]
CLEAN-UP        [/]     -- Time: [ 0.83s ]

Some duplicated files failed to move
Press Enter to exit...
```

## Unsort

```
--- UNSORTING ---

PARSING FILES   [/]     -- Time: [ 2.88s ] -- Files: [ 19483 ] -- Folders: [ 3532 ] --
MOVING FILES    [/]     -- Time: [ 14.89s ]
CLEAN-UP        [/]     -- Time: [ 0.61s ]

Some duplicated files failed to move
Press Enter to exit...
```

## Filter and Copy

```
Enter the file extension\s the you wish to filter: 
png jpeg

-- FILTERING --

PARSING FILES   [/]      -- Time: [ 3.28s ] -- Files: [ 26606 ]
COPYING FILES   [/]     -- Time: [ 3.94s ]
CLEAN-UP        [/]     -- Time: [ 0.46s ]

Some files with duplicated names might not be moved
Press Enter to exit...

```

## Filter and Move

```
Enter the file extension\s the you wish to filter: 
png

-- FILTERING --

PARSING FILES   [/]      -- Time: [ 2.82s ] -- Files: [ 17848 ]
MOVING FILES    [/]     -- Time: [ 7.48s ]
CLEAN-UP        [/]     -- Time: [ 0.43s ]

Some duplicated files failed to be copy/moved
Press Enter to exit...
```

## Help

```
-- HELP --
1. [ sort ]:
        -- sorts the files inside the chosen directory (not including ones inside folders).

2. [ unsort ]:
        -- unsorts all files and folders inside the chosen directory.

3. [ filter & copy ]:
        -- filters all files with the specified extension and creates a copy in a separate directory.

4. [ filter & move ]:
        -- filters all files with the specified extension and moves it into a separated directory.
```

# Installation

```bash
git clone https://github.com/theMyle/goFileSorter.git
cd goFileSorter
go install .
gofilesorter
```
you can now launch the app by simply typing `gofilesorter` in the commandline 

# Uninstall

```bash
# Windows
## PowerShell
rm $env:GOPATH/bin/gofilesorter.exe
## CMD
rm %GOPATH%/bin/gofilesorter.exe

# Linux
rm $GOPATH/bin/gofilesorter
```

# Notes

- if a file with same name already exists in the destination path, the file will not be moved and will therfore remain. 
