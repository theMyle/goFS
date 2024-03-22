# Go File Sorter

A simple CLI app for sorting files

## Features

- **SORT**: Sorts top level files in the chosen directory (not including
  directories).
- **UNSORT**: Unsorts all files in the chosen directory (recursively).
- **FILTER**: Filters files with the specified file extensions in the chosen
  directory.
  - _filter & move_ - moves files.
  - _filter & copy_ - copies files.

# Usage

```
PS C:\Users\jangk\Documents\Git\goFileSorter> go run .
<____> Go File Sorter v1.0 <____>

Select an option: (press Q to quit)
        1. sort
        2. unsort
        3. filter & copy
        4. filter & move
        5. help
>>
```

## Sort Sample

```
Do you really wish to (SORT) this directory? (y/n)
C:\Users\jangk\Downloads: y

--- SORTING ---

PARSING FILES   [/]     -- Files: [ 34664 ]
MOVING FILES    [/]
Finished        [/]

Total execution time: [ 28.6192154s ]
Press Enter to exit...
```

## Unsort Sample

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
Do you really wish to (FILTER & MOVE) this directory? (y/n)
C:\Users\jangk\Downloads: y

Enter the file extension\s the you wish to filter:
dat

-- FILTERING --

PARSING FILES   [/]     -- Files: [ 9969 ] -- Time: [ 216.5309ms ]
MOVING FILES    [/]     -- Time: [ 7.9612256s ]
Finished        [/]

Total execution time: [ 7.9612256s ]
Press Enter to exit...
```

## Filter and Copy

```
Do you really wish to (FILTER & COPY) this directory? (y/n)
C:\Users\jangk\Downloads: y

Enter the file extension\s the you wish to filter:
sab

-- FILTERING --

PARSING FILES   [/]     -- Files: [ 22105 ] -- Time: [ 175.237ms ]
COPYING FILES   [/]     -- Time: [ 6m2.7658263s ]
Finished        [/]

Total execution time: [ 6m2.7663698s ] 
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

- Filter and Copy it the slowest operation (I might remove it).
- Unsort will literally unsort everything so be careful.
