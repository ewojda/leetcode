package main

func main() {}

func isMatch(target, pattern string) bool {
	if len(pattern) == 0 && len(target) == 0 {
		return true
	} else if len(pattern) == 0 {
		return false
	}
	patternChar, patternMod, remainingTokens := nextPatternToken(pattern)
	for _, targetRemaining := range remainingAfterMatch(target, patternChar, patternMod) {
		if isMatch(targetRemaining, remainingTokens) {
			return true
		}
	}
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
