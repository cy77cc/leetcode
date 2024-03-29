package main

import "fmt"

func findMedianSortedArraysv1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	result := make([]int, 0)
	leftindex := 0
	rightindex := 0
	for leftindex < len(nums1) && rightindex < len(nums2) {
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
	if rightindex < len(nums2) {
		result = append(result, nums2[rightindex:]...)
	}
	if leftindex < len(nums1) {
		result = append(result, nums1[leftindex:]...)
	}
	fmt.Println(result)

	if length == 0 {
		return float64(0)
	}
	if length == 1 {
		return float64(result[0])
	}

	if length%2 == 0 {
		return float64(result[length/2]+result[length/2-1]) / 2
	} else {
		return float64(result[length/2])
	}
}

func findMedianSortedArraysv2(nums1 []int, nums2 []int) float64 {
	mid := (len(nums1) + len(nums2))/2
	if mid == 0 {
		if len(nums1) == 0 && len(nums2) == 0 {
			return 0
		} else if len(nums1) == 0 {
			return float64(nums2[0])
		} else {
			return float64(nums1[0])
		}
	} else {

	}
}

func main() {
	fmt.Println(findMedianSortedArraysv2([]int{1}, []int{2}))
}
