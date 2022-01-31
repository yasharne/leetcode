func thirdMax(nums []int) int {
	first := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > first {
			first = nums[i]
		}
	}

	if len(nums) < 3 {
		return first
	}

	nums2 := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] < first {
			nums2 = append(nums2, nums[i])
		}
	}

	if len(nums2) < 2 {
		return first
	}
	second := nums2[0]

	for i := 1; i < len(nums2); i++ {
		if nums2[i] > second {
			second = nums2[i]
		}
	}

	nums3 := make([]int, 0)

	for i := 0; i < len(nums2); i++ {
		if nums2[i] < second {
			nums3 = append(nums3, nums2[i])
		}
	}

	if len(nums3) < 1 {
		return first
	}
	third := nums3[0]

	for i := 1; i < len(nums3); i++ {
		if nums3[i] > third {
			third = nums3[i]
		}
	}

	return third
}