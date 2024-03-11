package internal

func MakeUniqueMap(fileNames []string) map[string]bool {
	uniqueMap := make(map[string]bool)

	for _, file := range fileNames {
		_, ok := uniqueMap[file]
		if ok {
			// Duplicated
			uniqueMap[file] = false
		} else {
			// Unique
			uniqueMap[file] = true
		}
	}

	return uniqueMap
}

func FilterDuplicated(mp map[string]bool) (unique []string, duplicated []string) {
	uniqueFile := make([]string, 0)
	duplicatedFile := make([]string, 0)

	for fileName, uniq := range mp {
		if uniq {
			uniqueFile = append(uniqueFile, fileName)
		} else {
			duplicatedFile = append(duplicatedFile, fileName)
		}
	}

	return uniqueFile, duplicatedFile
}

func IsUnique(name string, mp map[string]bool) bool {
	return mp[name]
}
