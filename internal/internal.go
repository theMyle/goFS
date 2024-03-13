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

func IsUnique(name string, mp map[string]bool) bool {
	return mp[name]
}
