package main

func twoSum2(nums []int, target int) []int {
	for left := 0; left < len(nums)-1; left++ {
		targetRight := target - nums[left]
		for right := left + 1; right < len(nums); right++ {
			if nums[right] == targetRight {
				return []int{left, right}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for left := 0; left < len(nums)-1; left++ {
		for right := left + 1; right < len(nums); right++ {
			if nums[left]+nums[right] == target {
				return []int{left, right}
			}
		}
	}
	return nil
}

func main() {}
