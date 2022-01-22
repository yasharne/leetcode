func removeDuplicates(nums []int) int {
	if len(nums) > 0 {
		i := 0
		for {
			if i+1 < len(nums) && nums[i+1] == nums[i] {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				i++
			}
			if i == len(nums) {
				break
			}
		}
	}
	return len(nums)
}