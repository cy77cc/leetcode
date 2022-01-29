package main

import "fmt"


// 垃圾代码
func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n % 2 != 0 {
		return false
	}
	i := 1
	res := 0
	for {
		res = 1 << i
		if res == n {
			return true
		}
		if res > n {
			return false
		}
		i++
	}
}

func main() {
	fmt.Println(isPowerOfTwo(-16))
}
