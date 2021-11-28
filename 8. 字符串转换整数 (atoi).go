package main

import (
	"fmt"
	"strconv"
)

func myAtoi(s string) int {
	bytes := []byte(s)
	result := make([]byte, 0)
	var flag bool = false
	if len(bytes) == 0 {
		return 0
	}
	for i := 0; i < len(bytes); i++ {
		if (bytes[i] == '-' || bytes[i] == '+') && len(result) == 0 {
			if len(bytes) == 1 {
				continue
			}
			if bytes[i+1] >= '0' && bytes[i+1] <= '9' {
				result = append(result, bytes[i])
				continue
			}
		}
		flag = (bytes[i] >= 'A' && bytes[i] <= 'Z') || (bytes[i] >= 'a' && bytes[i] <= 'z') || (bytes[i] == '+') || (bytes[i] == '-') || (bytes[i] == '.')
		if len(result) != 0 && (flag || bytes[i] == ' ') {
			break
		}

		if len(result) == 0 && flag {
			break
		}
		if bytes[i] >= '0' && bytes[i] <= '9' {
			result = append(result, bytes[i])
		}
	}
	atoi, _ := strconv.Atoi(string(result))
	if atoi < -2147483648 {
		return -2147483648
	} else if atoi > 2147483647 {
		return 2147483647
	} else {
		return atoi
	}
}

func equal(e byte) bool {
	return (e >= 'A' && e <= 'Z') || (e >= 'a' && e <= 'z') || (e == '+') || (e == '-') || (e == '.')
}

func myAtoi2(s string) int {
	bytes := []byte(s)
	result := make([]byte,0)
	if len(bytes) == 0 {
		return 0
	}
	if len(bytes) == 1 {
		if equal(bytes[0]) {
			return 0
		} else if bytes[0] == ' ' {
			return 0
		} else {
			atoi, _ := strconv.Atoi(string(bytes[0]))
			return atoi
		}
	}
	var pre, next byte
	for i := 0; i < len(bytes) - 1; i += 2 {
		pre = bytes[i]
		next = bytes[i+1]
		fmt.Println(pre)
		fmt.Println(next)
		if (pre == '+' || pre == '-') && (next >= '0' && next <= '9') {
			result = append(result, pre, next)
			continue
		}
		if len(result) == 0 && (equal(pre)) {
			return 0
		}else if len(result) != 0 && (equal(pre) && equal(next)) {
			break
		} else if pre >= '0' && pre <= '9' {
			if equal(next) || next == ' ' {
				result = append(result, pre)
				break
			} else  {
				result = append(result, pre, next)
			}
		}
	}
	if len(bytes) % 2 != 0 {
		if bytes[len(bytes) - 1] >= '0' || bytes[len(bytes) - 1] <= '9' {
			result = append(result, bytes[len(bytes) - 1])
		}
	}
	atoi, _ := strconv.Atoi(string(result))
	return atoi
}

func main() {
	//fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi2("+-1"))
}
