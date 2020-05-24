package main

import (
	"math"
)

func main() {}

func myAtoi(str string) int {
	sign := 0
loop:
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ' ':
			if sign != 0 {
				return 0
			}
		case '+':
			if sign != 0 {
				return 0
			}
			sign = 1
		case '-':
			if sign != 0 {
				return 0
			}
			sign = -1
		case '0':
			if i == len(str)-1 {
				return 0
			}
			if sign == 0 {
				sign = 1
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if sign == 0 {
				sign = 1
			}
			str = str[i:]
			break loop
		default:
			return 0
		}
	}
	mult := int64(1)
	result := int64(0)
	for i := len(str) - 1; i >= 0; i-- {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result += (int64(str[i]) - int64('0')) * mult
			if result > math.MaxInt32 || result < math.MinInt32 || mult > math.MaxInt32 || mult < math.MinInt32 {
				switch sign {
				case 1:
					return math.MaxInt32
				case -1:
					return math.MinInt32
				}
			}
			mult *= 10
		default:
			mult = 1
			result = 0
		}
	}
	return int(result) * sign
}
