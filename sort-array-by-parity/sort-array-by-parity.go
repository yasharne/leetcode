func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return nums
}