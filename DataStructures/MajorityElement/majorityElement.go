package majorityElement

import "math"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	countMap := make(map[int]int)
	majorityElm := math.Floor(float64(len(nums) / 2))
	for _, num := range nums {
		if x, ok := countMap[num]; ok {
			if float64(x+1) > majorityElm {
				return num
			} else {
				countMap[num] = x + 1
			}
		} else {
			countMap[num] = 1
		}
	}

	return 0
}
