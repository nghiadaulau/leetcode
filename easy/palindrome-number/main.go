package palindrome_number

func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	return n == reverseNumber(n)
}
