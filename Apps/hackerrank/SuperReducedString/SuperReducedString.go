package SuperReducedString

func SuperReducedString(s string) string {
	current := removeDups(s)
	done := len(current) == len(s)
	for !done {
		temp := current
		current = removeDups(temp)
		done = len(current) == len(temp)
	}

	return current
}

func removeDups(s string) string {
	tempRes := ""
	for i := 0; i < len(s); {
		if i+1 == len(s) {
			tempRes = tempRes + string(s[i])
			break
		}
		if s[i] == s[i+1] {
			i = i + 2
		} else {
			tempRes = tempRes + string(s[i])
			i++
		}
	}
	// tempRes = tempRes + string(s[len(s)-1])
	return tempRes
}
