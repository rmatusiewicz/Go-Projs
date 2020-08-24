package main

func main() {}

func NumIdenticalPairs(nums []int) int {
	elementMap := make(map[int]int)
	for _, num := range nums {
		indexcount := elementMap[num]
		if indexcount != 0 {
			elementMap[num] = indexcount + 1
		} else {
			elementMap[num] = 1
		}
	}

	var countMatchingPairs int
	for key := range elementMap {
		x := elementMap[key]
		tempSum := (x*x - x) / 2
		countMatchingPairs = countMatchingPairs + tempSum
	}

	return countMatchingPairs
}
