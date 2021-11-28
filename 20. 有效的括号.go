package main

import (
	"fmt"
)

/*
	1. 思路两个符号之间相差偶数个位置
	2. 正确匹配
	3. 匹配一个出来就让flag变为true，经历
*/

func isValid(s string) bool {
	stack := make([]byte, 0)
	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 {
			if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[][][]()()"))
}
