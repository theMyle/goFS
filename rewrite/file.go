package internal

import "path/filepath"

// Contains data for file storage manipulation
type File struct {
	name         string
	extension    string
	absolutePath string
	parentDir    string
}

// Initializes a File struct
func NewFile(fileName string, absPath string) File {
	return File{
		name:         fileName,
		extension:    getExtension(fileName),
		absolutePath: absPath,
		parentDir:    getParentDir(absPath),
	}
}

// Returns file name
func (f *File) Name() string {
	return f.name
}

// Returns file extension
func (f *File) Extension() string {
	return f.extension
}

// Returns file absoulute path
func (f *File) AbsPath() string {
	return f.absolutePath
}

// Returns file parent directory
func (f *File) ParentDir() string {
	return f.parentDir
}

// Backwards extension getter function:
//
//	returns nil when no valid extension is found
func getExtension(file string) string {
	len := len(file)
	for len != 0 {
		if file[len-1] == '.' {
			if len > 1 {
				if file[len-2] == byte(filepath.Separator) {
					return ""
				}
				return file[len:]
			}
		}
		len--
	}
	return ""
}

// Returns parent directory if valid
//
//	returns nil if invalid
func getParentDir(directory string) string {
	len := len(directory)
	for len != 0 {
		if len > 1 {
			if directory[len-1] == '/' || directory[len-1] == '\\' {
				if len > 2 {
					if directory[len-2] == '/' || directory[len-2] == '\\' {
						return ""
					}
				}
				return directory[:len]
			}
		}
		len--
	}
	return ""
}
