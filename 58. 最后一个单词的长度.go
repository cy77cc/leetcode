package main

import "fmt"

func lengthOfLastWord(s string) int {
	if len(s) == 1 {
		return 1
	}
	left, right := len(s) - 1, len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			left--
		} else {
			left--
			right--
		}
		if left != right && (left < 0 ||s[left] == ' ' ) {
			break
		}
	}
	return right - left

}

func main() {
	fmt.Println('a' + 32)
}
