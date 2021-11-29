package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for ; left < right; {
		if nums[left] == val {
			nums[left], nums[right-1] = nums[right-1], nums[left]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 4, 1, 5, 2}, 1))
}
