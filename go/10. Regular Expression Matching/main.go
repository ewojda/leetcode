package main

func main() {}

func isMatch(target, pattern string) bool {
	memo := map[[2]int]bool{}
	return isMatch0(target, pattern, memo)
}

func isMatch0(target, pattern string, memo map[[2]int]bool) bool {
	if memoResult, ok := memo[[2]int{len(target), len(pattern)}]; ok {
		return memoResult
	}

	if len(pattern) == 0 && len(target) == 0 {
		if len(target) == 0 {
			memo[[2]int{len(target), len(pattern)}] = true
			return true
		} else {
			memo[[2]int{len(target), len(pattern)}] = false
			return false
		}
	}
	patternChar, patternMod, remainingTokens := nextPatternToken(pattern)
	for _, targetRemaining := range remainingAfterMatch(target, patternChar, patternMod) {
		if isMatch0(targetRemaining, remainingTokens, memo) {
			memo[[2]int{len(target), len(pattern)}] = true
			return true
		}
	}
	memo[[2]int{len(target), len(pattern)}] = false
	return false
}

func nextPatternToken(pattern string) (byte, byte, string) {
	var char, mod byte
	if len(pattern) > 0 && isValidCharacter(pattern[0]) {
		char = pattern[0]
		pattern = pattern[1:]
		if len(pattern) > 0 && isValidModifier(pattern[0]) {
			mod = pattern[0]
			pattern = pattern[1:]
		}
	}
	return char, mod, pattern
}

func remainingAfterMatch(target string, patternChar, patternMod byte) []string {
	if !isValidCharacter(patternChar) {
		return []string{}
	}
	switch patternMod {
	case 0:
		if len(target) == 0 {
			return []string{}
		}
		if charMatches(patternChar, target[0]) {
			return []string{target[1:]}
		} else {
			return []string{}
		}
	case '*':
		result := []string{target}
		for {
			if len(target) == 0 {
				return result
			}
			if charMatches(patternChar, target[0]) {
				target = target[1:]
				result = append(result, target)
			} else {
				return result
			}
		}
	default:
		panic("Unrecognized modifier: '" + string(patternMod) + "'")
	}
}

func charMatches(patternChar, targetChar byte) bool {
	return patternChar == '.' || patternChar == targetChar
}

func isValidCharacter(char byte) bool {
	return char >= 'a' && char <= 'z' || char == '.'
}

func isValidModifier(char byte) bool {
	return char == '*' || char == 0
}
