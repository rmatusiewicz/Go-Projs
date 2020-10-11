package basicCalculator

import (
	"regexp"
	"strconv"
)

var ignoredChars map[string]string

func BasicCalculator(s string) int {
	var operant string
	var currNum int
	var numBuild string
	// ignoredChars := []string{"(", ")", " "}
	ignoredChars = make(map[string]string)
	ignoredChars["("] = ""
	ignoredChars[")"] = ""
	ignoredChars[" "] = ""
	res, startFrom := initRes(s)

	for i := startFrom; i < len(s); i++ {
		strc := string(s[i])

		if inIgnoredCharacters(strc) {
			continue
		} else if strc == "+" || strc == "-" {
			//Have first num
			if operant == "" {
				operant = strc
				numBuild = ""
				continue
			}

			//We now have the second num
			currNum, _ = strconv.Atoi(numBuild)
			res = operate(operant, res, currNum)
			operant = strc
			numBuild = ""
		} else { //numeric
			numBuild = numBuild + strc
		}
	}

	lastNum, _ := strconv.Atoi(numBuild)
	res = operate(operant, res, lastNum)

	return res
}

func operate(op string, arg1 int, arg2 int) int {
	if op == "-" {
		return arg1 - arg2
	}
	//+
	ret := arg1 + arg2
	return ret
}

func initRes(s string) (int, int) {
	re := regexp.MustCompile(`[0-9]+`)
	indexSlc := re.FindStringIndex(s)
	a, b := indexSlc[0], indexSlc[1]
	num, _ := strconv.Atoi(s[a:b])
	return num, b
}

func inIgnoredCharacters(checkStr string) bool {
	if _, yes := ignoredChars[checkStr]; yes {
		return true
	}
	return false
}
