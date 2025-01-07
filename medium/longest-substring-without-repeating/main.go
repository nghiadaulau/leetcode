package longest_substring_without_repeating

func contains(arr []rune, value rune) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	var runes []rune
	maxLength := 0

	for _, char := range s {
		for contains(runes, char) {
			runes = runes[1:]
		}

		runes = append(runes, char)

		if len(runes) > maxLength {
			maxLength = len(runes)
		}
	}

	return maxLength
}
