package main


//[0,1,0,3,12]
/*
	left = 0
	right = 0

	1. right < n true
		right = 1
	2. right < n true
		[1, 0, 0, 3, 12]
		right = 2
		left = 1
	3. right < n true
		right = 3
	4. right < n true
		[1, 3, 0, 0, 12]
		right = 4
		left = 2
	5. right < n true
		[1, 3, 12, 0, 0]
		right = 5
		left = 3

*/
func moveZeroes(nums []int)  {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {

}
