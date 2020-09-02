package maxVowels

import "fmt"

func main() {

	// fmt.Println(isVowel("a"))
	//  fmt.Println(isVowel("j"))
	input1 := "abcdeeeefg"
	fmt.Printf("\nTest 1: %s\n", input1)
	test1 := maxVowels(input1, 4)
	// fmt.Println(test1)
	if test1 != 4 {
		fmt.Printf("\nERROR: Expected 4 but got %d", test1)
	} else {
		fmt.Printf("PASS\n")
	}

	input2 := "ab"
	fmt.Printf("\nTest 2: %s\n", input2)
	test2 := maxVowels(input2, 4)
	if test2 != 1 {
		fmt.Printf("\nERROR: Expected 1 but got %d\n", test2)
	} else {
		fmt.Printf("PASS\n")

	}

	input3 := "abababab"
	fmt.Printf("\nTest 3: %s\n", input3)
	test3 := maxVowels(input3, 4)
	if test3 != 2 {
		fmt.Printf("\nERROR: Expected 1 but got %d\n", test3)
	} else {
		fmt.Printf("PASS\n")

	}

	input4 := "aaaaaa"
	fmt.Printf("\nTest 4: %s", input4)
	test4 := maxVowels(input4, 4)
	if test4 != 4 {
		fmt.Printf("\nERROR: Expected 4 but got %d\n", test4)
	} else {
		fmt.Printf("PASS\n")

	}

	fmt.Printf("\nTest 5\n")
	test5 := maxVowels("weallloveyou", 7)
	if test5 != 4 {
		fmt.Printf("\nERROR: Expected 4 but got %d\n", test5)
	} else {
		fmt.Printf("PASS\n")
	}

	input6 := "qempburycnhrvvccr"
	fmt.Printf("\nTest 6: %s\n", input6)
	test6 := maxVowels(input6, 13)
	if test6 != 2 {
		fmt.Printf("\nERROR: Expected 2 but got %d\n", test6)
	} else {
		fmt.Printf("PASS\n")
	}

}

func maxVowels(s string, k int) int {
	bWin := 0
	var tWin int
	if k > len(s) {
		tWin = len(s) - 1
	} else {
		tWin = k - 1
	}

	count := 0
	for i := 0; i <= tWin; i++ {
		if isVowel(string(s[i])) {
			count++
		}
	}
	if k > len(s) {
		return count
	}

	fmt.Printf("\nCount: %d\n", count)

	max := count
	for tWin < len(s) {
		fmt.Printf("\n%d: %s  %d: %s\n", bWin, string(s[bWin]), tWin, string(s[tWin]))

		if isVowel(string(s[bWin])) {
			count--
		}
		bWin++
		tWin++

		if tWin < len(s) && isVowel(string(s[tWin])) {
			count++
			if count > max {
				max = count
			}
		}
		fmt.Printf("\nCount %d\n", count)
	}

	return max
}

func isVowel(letter string) bool {
	if letter == "a" ||
		letter == "e" ||
		letter == "i" ||
		letter == "o" ||
		letter == "u" {
		return true
	}
	return false
}
