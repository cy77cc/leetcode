package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	result := make([]int, 0)
	leftindex := 0
	rightindex := 0
	for leftindex < m && rightindex < n {
		if nums1[leftindex] < nums2[rightindex] {
			result = append(result, nums1[leftindex])
			leftindex++
		} else if nums1[leftindex] > nums2[rightindex] {
			result = append(result, nums2[rightindex])
			rightindex++
		} else {
			result = append(result, nums1[leftindex])
			result = append(result, nums2[rightindex])
			leftindex++
			rightindex++
		}
	}
	if rightindex < n {
		result = append(result, nums2[rightindex:n]...)
	}
	if leftindex < m {
		result = append(result, nums1[leftindex:m]...)
	}
	copy(nums1, result)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr1, 4, arr2, 3)
	fmt.Println(arr1)
}
