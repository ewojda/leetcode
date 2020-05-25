package main

func main() {}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longest := ""
	longestPalindromeCenteredAt := func(left, right int) string {
		if s[left] != s[right] {
			return ""
		}
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				return s[left+1 : right]
			}
			left--
			right++
		}
	}
	setIfNewLongest := func(left, right int) {
		if newLongest := longestPalindromeCenteredAt(left, right); len(newLongest) > len(longest) {
			longest = newLongest
		}
	}
	for i := 0; i < len(s)-1; i++ {
		setIfNewLongest(i, i+1)
	}
	for i := 0; i < len(s)-2; i++ {
		setIfNewLongest(i, i+2)
	}
	if longest == "" {
		return s[:1]
	}
	return longest
}
