package main

import "fmt"

func strStr(haystack string, needle string) int {
	index := -1
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}
	if lenHaystack == 0{
		return -1
	}
	for i := 0; i <= lenHaystack - lenNeedle; i++ {
		fmt.Println(haystack[i:i+lenNeedle])
		if haystack[i:i+lenNeedle] == needle {
			index = i
			break
		} else {

		}
	}
	return index
}

func main() {
	fmt.Println(strStr("", "bba"))
}
