package main

func main() {}

func CountSegs(s string) int {
	// s = strings.TrimSpace(s)
	// if len(s) == 0 {
	// 	return 0
	// }

	segCount := 0
	var seg string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			seg = seg + string(s[i])
		} else { //Space
			if seg != "" {
				segCount++
				seg = ""
			}
		}
	}
	if seg != "" {
		segCount++
	}

	return segCount
}
