package validPalindrome

import (
	"log"
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, "")

	var i1 int
	if len(s)%2 == 0 {
		i1 = len(s) / 2
	} else {
		i1 = (len(s) - 1) / 2
	}

	j := len(s) - 1
	for i := 0; i < i1; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}
