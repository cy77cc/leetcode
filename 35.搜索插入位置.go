package main

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target < nums[0] {
			nums = append([]int{target}, nums...)
			return  0
		}
		if nums[i] == target {
			return i
		}
		if nums[i] > target && nums[i-1] < target {
			nums = append(append(nums[:i], target), nums[i:]...)
			return i
		}
		if nums[i] < target && i == len(nums)-1 {
			nums = append(nums, target)
			return len(nums)-1
		}
	}
	return 0
}

func main() {

}
