package main

func main() {}

func Shuffle(nums []int, n int) []int {
	//nums = [x1,x2,...,xn,y1,y2,...,yn]
	//[x1,y1,x2,y2,...,xn,yn]
	//n -> starting index of y's

	var shuffleDeck []int
	for i := 0; i < n; i++ {
		//add x
		shuffleDeck = append(shuffleDeck, nums[i])
		//add y
		shuffleDeck = append(shuffleDeck, nums[n+i])
	}
	return shuffleDeck
}
