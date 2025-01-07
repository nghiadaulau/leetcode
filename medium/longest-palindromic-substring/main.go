package longest_palindromic_substring

func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	longest := ""

	for i := 0; i < len(s); i++ {
		odd := expandAroundCenter(s, i, i)
		even := expandAroundCenter(s, i, i+1)

		if len(odd) > len(longest) {
			longest = odd
		}
		if len(even) > len(longest) {
			longest = even
		}
	}

	return longest
}
