package longest_common_prefix

func shortestString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]

	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	return shortest
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	longPrefix := ""
	shortWord := shortestString(strs)
	for index, value := range shortWord {
		isPrefix := false
		prefix := ""
		for _, value2 := range strs {
			if string(value2[index]) != string(value) {
				isPrefix = false
				prefix = ""
				break
			} else {
				isPrefix = true
				prefix = string(value2[index])
			}
		}
		if !isPrefix {
			return longPrefix
		}
		if isPrefix {
			longPrefix += prefix
		}
	}
	return longPrefix
}
