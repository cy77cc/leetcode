package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n - 1
	for pos := n-1; pos >= 0; pos-- {
		if v, m := nums[i] * nums[i], nums[j] * nums[j]; v > m {
			ans[pos] = v
			i++
		} else if v < m {
			ans[pos] = m
			j--
		} else {
			ans[pos] = v
			pos--
			i++
			ans[pos] = m
			j--
		}
	}
	return ans
}

func main() {

}
