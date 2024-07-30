# Go File Sorter

A simple CLI app for sorting files

## Features

- `SORT`: Sorts the entire content of the selected directory.
- `UNSORT`: Unsorts all files in the chosen directory (Puts them outside).
- `FILTER`: Filters files with the specified file extensions in the chosen
  directory.
  - `_filter & move_` - moves files.
  - `_filter & copy_` - copies files.

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
Do you really wish to sort this directory?
[C:\Users\jangk\Downloads\test] (y/n): y

--- SORTING ---

PARSING FILES   [/]     -- Time: [ 7.12s ] -- Files: [ 19483 ]
MOVING FILES    [/]     -- Time: [ 14.90s ]
CLEAN-UP        [/]     -- Time: [ 1.18s ]

Some duplicated files failed to move
Press Enter to exit...
```

## Unsort

```
Do you really wish to (UNSORT) this directory? (y/n)
C:\Users\jangk\Downloads: y

--- UNSORTING ---

PARSING FILES   [/]     -- Files: [ 34712 ] -- Folders: [ 229 ] -- Time: [ 146.5766ms ] --
MOVING FILES    [/]     -- Time: [ 39.5003085s ]
CLEANUP         [/]
Finished        [/]

Total execution time: [ 39.653121s ]
Press Enter to exit...
```

## Filter and Move

```
Implementation not yet polished
```

## Filter and Copy

```
Implementation not yet polished
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
