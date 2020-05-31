package main

func main() {}

func maxArea(height []int) int {
	max := 0
outer:
	for left := 0; left < len(height)-1; left++ {
		for right := len(height) - 1; right > left; right-- {
			lh, rh := height[left], height[right]
			lower := lh
			if rh < lower {
				lower = rh
			}
			area := lower * (right - left)
			if area > max {
				max = area
			}
			if rh >= lh {
				continue outer
			}
		}
	}
	return max
}
