package main

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	longest := 1
	characterInSubstring := [256]bool{}
	hasNoRepeatingCharacters := func(substring string) bool {
		for i := range characterInSubstring {
			characterInSubstring[i] = false
		}
		for i := range substring {
			if characterInSubstring[substring[i]] {
				return false
			} else {
				characterInSubstring[substring[i]] = true
			}
		}
		return true
	}

outer:
	for start := 0; start < len(s); start++ {
		for end := start + longest; end <= len(s); end++ {
			substring := s[start:end]
			if hasNoRepeatingCharacters(substring) {
				if len(substring) > longest {
					longest = len(substring)
				}
			} else {
				continue outer
			}
		}
	}
	return longest
}

func main() {}
