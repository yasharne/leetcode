func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && i < len(nums)-1 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}