package main

func main() {}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var most []bool
	maxFound := 0

	for _, kid := range candies {
		if kid > maxFound {
			maxFound = kid
		}
	}

	for _, child := range candies {
		if child+extraCandies < maxFound {
			most = append(most, false)
		} else {
			most = append(most, true)
		}
	}

	return most
}
