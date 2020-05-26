package main

import (
	"math"
)

func main() {}

func isPalindrome(x int) bool {
	return x >= 0 && x == reverse(x)
}

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	result := 0
	for digitLow, digitHigh := 1, 10; digitLow <= x; digitLow, digitHigh = digitLow*10, digitHigh*10 {
		digit := x / digitLow
		digit %= 10
		if mulOverflows(result, 10, sign) {
			return 0
		}
		result *= 10
		if sumOverflows(result, digit, sign) {
			return 0
		}
		result += digit
	}
	return result * sign
}

func sumOverflows(a, b, sign int) bool {
	max := math.MaxInt32
	if sign == -1 {
		max = -math.MinInt32
	}
	return a > max-b
}

func mulOverflows(a, b, sign int) bool {
	max := math.MaxInt32
	if sign == -1 {
		max = -math.MinInt32
	}
	if b == 0 {
		return false
	}
	return a > max/b
}
