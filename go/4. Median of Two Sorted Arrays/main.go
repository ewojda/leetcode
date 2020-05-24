package main

func main() {}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	NUMS1, NUMS2 := true, false

	lentotal := len(nums1) + len(nums2)
	medianI := lentotal / 2
	i, i1, i2 := 0, 0, 0
	median, lastMedian := 0, 0

	next := func(nums []int, i int) (int, bool) {
		if len(nums) > i {
			return nums[i], true
		} else {
			return -1, false
		}
	}

	pushMedian := func(newMedian int, whichNums bool) {
		lastMedian = median
		median = newMedian
		if whichNums == NUMS1 {
			i1++
		} else {
			i2++
		}
		i++
	}

	for i <= medianI {
		num1, num1ok := next(nums1, i1)
		num2, num2ok := next(nums2, i2)
		if num1ok && num2ok {
			if num1 < num2 {
				pushMedian(num1, NUMS1)
			} else {
				pushMedian(num2, NUMS2)
			}
		} else if num1ok && !num2ok {
			pushMedian(num1, NUMS1)
		} else if num2ok && !num1ok {
			pushMedian(num2, NUMS2)
		} else {
			panic("Shouldn't happen")
		}
	}
	if lentotal%2 != 0 {
		return float64(median)
	} else {
		return (float64(median) + float64(lastMedian)) / 2.0
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lentotal := len(nums1) + len(nums2)
	medianI, i, i1, i2 := lentotal/2, 0, 0, 0
	nums := make([]int, 0, lentotal)

	if i1 >= len(nums1) {
		
	}
}
