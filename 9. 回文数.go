package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	itoa := strconv.Itoa(x)
	left, right := 0, len(itoa) - 1
	flag := true
	if itoa[left] == '-' {
		flag = false
	} else {
		for ; left < right; {
			if itoa[left] != itoa[right] {
				flag = false
			}
			left++
			right--
		}
	}
	return flag
}

func main() {
	fmt.Println(isPalindrome(1001001))
}
