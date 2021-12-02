package main

import "fmt"

func reverseWords(s string) string {
	left, right, n := 0, 0, 0
	bytes := []byte(s)
	for {
		if left >= len(bytes)-1 {
			break
		}
		if right == len(bytes) || bytes[right] == ' ' {
			n = right
			right--
			for left <= right {
				bytes[left], bytes[right] = bytes[right], bytes[left]
				left++
				right--
			}
			left = n + 1
			right = n + 1
		} else {
			right++
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(reverseWords("content"))
}
