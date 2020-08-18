package main

import "fmt"

func main() {
	x := isValid("}()")
	fmt.Print("\n", x)
}

func isValid(s string) bool {
	fmt.Println("TEST:", s)
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	if s[0] == ')' || s[0] == '}' || s[0] == ']' {
		return false
	}

	var stack []rune

	for _, char := range s {
		fmt.Printf("\nStack: %v", stack)
		if char == '(' {
			stack = append(stack, '(')
			continue
		}
		if char == '{' {
			stack = append(stack, '{')
			continue
		}
		if char == '[' {
			stack = append(stack, '[')
			continue
		}

		if char == ')' {
			if stack[len(stack)-1] == '(' {
				stack = pop(stack)
				continue
			} else {
				return false
			}
		}
		if char == '}' {
			if stack[len(stack)-1] == '{' {
				stack = pop(stack)
				continue
			} else {
				return false
			}
		}
		if char == ']' {
			if stack[len(stack)-1] == '[' {
				stack = pop(stack)
				continue
			} else {
				return false
			}
		}
	}

	fmt.Printf("\nStack end: %v\n", stack)
	if len(stack) > 0 {
		return false
	}
	return true
}

func pop(s []rune) []rune {
	fmt.Print("\nPOP\n")
	if len(s)-1 > 0 {
		return s[:len(s)-1]
	}
	return []rune{}
}
