package goutils

func RemoveExcludedFromSlice(slice []string, exclude []string) []string {
	var excludedSlice []string
	for _, item := range slice {
		valid := true
		for _, excludedItem := range exclude {
			if excludedItem == item {
				valid = false
				break
			}
		}
		if valid {
			excludedSlice = append(excludedSlice, item)
		}
	}
	return excludedSlice
}

func KeepIncludedInSlice(slice []string, include []string) []string {
	var includedSlice []string
	for _, item := range slice {
		valid := false
		for _, includedItem := range include {
			if includedItem == item {
				valid = true
				break
			}
		}
		if valid {
			includedSlice = append(includedSlice, item)
		}
	}
	return includedSlice
}
