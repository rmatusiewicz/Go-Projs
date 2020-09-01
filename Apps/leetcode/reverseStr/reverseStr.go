package reverseStr

func ReverseStr(s string, k int) string {
	result := ""
	l := len(s)
	i := 0
	for i < l {

		if l-i >= 2*k {
			for j := k; j > 0; j-- {
				result += string(s[i+j-1])
			}
			i = i + k

			for j := 0; j < k; j++ {
				result += string(s[i+j])
			}
			i = i + k
			continue
		}

		if l-i >= k && l-i < 2*k {
			for j := k; j > 0; j-- {
				result += string(s[i+j-1])
			}
			i = i + k

			for i < l {
				result += string(s[i])
				i++
			}

			break
		}

		if l-i < k {
			//reverse all remaining
			for j := l - 1; j >= i; j-- {
				result += string(s[j])
			}
			break
		}
	}

	return result
}
