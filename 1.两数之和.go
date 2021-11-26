package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

func main() {
	arr := []int{2, 3, 4, 7, 2, 7, 9}
	fmt.Println(twoSum(arr, 9))

}
