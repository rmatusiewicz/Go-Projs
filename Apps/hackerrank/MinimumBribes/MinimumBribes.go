package MinimumBribes

func MinimumBribes(q []int32) int {
	sum := 0
	goodThrough := -1

	for goodThrough < len(q) {
		current := int(q[goodThrough+1])
		offBy := current - (goodThrough + 2) // +1 == actual val
		if offBy == 0 {
			goodThrough++
			if goodThrough == len(q)-2 {
				return sum
			}
			continue
		} else if offBy > 2 {
			sum = -1
			break
		} else {
			numBribes := swapRemaining(q, goodThrough+1)
			if numBribes == -1 {
				return -1
			}
			sum = sum + numBribes
		}
	}
	return sum
}

func swapRemaining(q []int32, i int) int {
	swaps := 0
	for i < len(q) {
		current := int(q[i])
		if current-(i+1) > 2 {
			return -1
		}

		next := int(q[i+1])
		if current > next {
			//swap
			q[i] = int32(next)
			q[i+1] = int32(current)
			swaps++
		}
		i++
		if i+1 >= len(q) {
			break
		}
	}
	return swaps
}
