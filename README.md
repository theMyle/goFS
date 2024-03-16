# Go File Sorter

A simple CLI app for sorting files

## Features

- **SORT**: Sorts top level files in the chosen directory (not including directories).
- **UNSORT**: Unsorts all files in the chosen directory (recursively).
- **FILTER**: Filters files with the specified file extensions in the chosen directory.
  - *filter & move* - moves files.
  - *filter & copy* - copies files.

## Usage
Run the executable or use "go run" command if you cloned the repo.
```
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

## Examples
