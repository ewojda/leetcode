package main

func main() {}

func maxArea(height []int) int {
	max := 0
	for left := 0; left < len(height)-1; left++ {
		for right := left + 1; right < len(height); right++ {
			lh, rh := height[left], height[right]
			lower := lh
			if rh < lower {
				lower = rh
			}
			area := lower * (right - left)
			if area > max {
				max = area
			}
		}
	}
	return max
}
