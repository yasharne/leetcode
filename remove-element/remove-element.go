func removeElement(nums []int, val int) int {
	if len(nums) > 0 {
		i := 0
		for {
			if nums[i] == val {
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
