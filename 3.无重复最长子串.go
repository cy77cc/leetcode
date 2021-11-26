package main

import "math"

func lengthOfLongestSubstring(s string) int {

	var last [128]int

	n := len(s)

	res := 0
	start := 0
	for i := 0; i < n; i++ {
		index := s[i]
		start = int(math.Max(float64(start), float64(last[index])))
		res = int(math.Max(float64(res), float64(i - start + 1)))
		last[index] = i+1
	}

	return res
}

func main() {

}
