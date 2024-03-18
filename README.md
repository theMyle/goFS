# Go File Sorter

A simple CLI app for sorting files

## Features

- **SORT**: Sorts top level files in the chosen directory (not including directories).
- **UNSORT**: Unsorts all files in the chosen directory (recursively).
- **FILTER**: Filters files with the specified file extensions in the chosen directory.
  - *filter & move* - moves files.
  - *filter & copy* - copies files.

## Usage
```
PS C:\Users\jangk\Documents\Programming\go\goFileSorter> go run .\cmd\main\GoSort.go
<____> Go File Sorter v1.0 <____>

Select an option: (press Q to quit)
        1. sort
        2. unsort
        3. filter & copy
        4. filter & move
        5. help
>> 1

Do you really wish to (SORT) this directory? (y/n)
C:\Users\jangk\Downloads: y

--- SORTING ---

PARSING FILES   [/]     -- Files: [ 10 ]
MOVING FILES    [/]
Finished        [/]

Total execution time: [ 8.3128ms ]
````

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

## Installation
```bash
git clone https://github.com/theMyle/goFileSorter.git
go install .\goFileSorter\cmd\main\GoSort.go

# Run by executing GoSort.exe
GoSort
```

## Uninstall
```bash
# Windows
rm $env:GOPATH/bin/GoSort

# Linux
rm $GOPATH/bin/GoSort
```

## Notes
- filter & copy seems to be the slowest.
- unsort will literally unsort everything so be careful.

## Issue/s
- error when file have no extension



