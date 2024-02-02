package folder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Folder struct {
	FolderName string
	FolderPath string
	Files      []string
	Subfolders []*Folder
}

func New_Folder(name string, path string) Folder {
	f := Folder{
		FolderName: name,
		FolderPath: path,
		Files:      []string{},
		Subfolders: []*Folder{},
	}

	return f
}

func Dir_Scan(f *Folder) []string {
	extensions := []string{}

	fp, fn := f.FolderPath, f.FolderName
	fname := strings.Split(fp, string(filepath.Separator))

	if fname[len(fname)-1] != fn {
		f.FolderPath = filepath.Join(fp, fn)
	}

	entries, err := os.ReadDir(f.FolderPath)
	if err != nil {
		fmt.Println("Error Scanning Directory")
		os.Exit(0)
	}

	for _,v := range entries {
		if v.IsDir() {
			path := filepath.Join(fp, v.Name())
			nf := New_Folder(v.Name(), path)	
			f.Subfolders = append(f.Subfolders, &nf)
			fmt.Printf("Dir: %s \t| Path: %s\n", path, v.Name())
			Dir_Scan(&nf)
		} else {
			fmt.Printf("File: %s \t| Path: %s\n", v.Name(), fp)
			f.Files = append(f.Files, v.Name())
			extensions = append(extensions, Get_Extension(v.Name()))
		}
	}

	return extensions
}

func Move_Items(f *Folder, root_path string) {
	rt := root_path

	for _,v := range f.Files {
		ext := Get_Extension(v)
		old_location := filepath.Join(f.FolderPath, v)
		new_location := filepath.Join(root_path, "Go_Sort", ext, v)

		// did not check if file exists
		fmt.Println(old_location, "->", new_location)
		os.Rename(old_location, new_location)
	}

	for _,j := range f.Subfolders {
		Move_Items(j, rt)
	}

	// Check if folder is empty
}

func Get_Extension(str string) string {
	split := strings.Split(str, ".")
	return split[len(split) - 1]
}