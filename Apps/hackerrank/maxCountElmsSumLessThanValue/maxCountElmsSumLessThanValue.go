package maxCountElmsSumLessThanValue

import "sort"

func maximumToys(prices []int32, k int32) int32 {
  // priceToCount := make(map[int32]int)
  // currentMinPrice := 0
  sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
  count := 0
  for _,p := range prices {
    if k > p {
      k = k - p
      count ++
    }else{
      break
    }
  }

  return int32(count)
}
