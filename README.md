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
>> 
````

## Examples
