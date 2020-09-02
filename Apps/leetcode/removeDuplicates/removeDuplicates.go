package removeDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	returnIndex := 1
	currentNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != currentNum {
			currentNum = nums[i]
			nums[returnIndex] = nums[i]
			returnIndex++
		}
	}

	return returnIndex
}
