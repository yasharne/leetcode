func sortArrayByParity(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			tmp := []int{nums[i]}
			tmp = append(tmp, nums[:i]...)
			tmp = append(tmp, nums[i+1:]...)
			_ = copy(nums, tmp)
		}
	}
	return nums
}