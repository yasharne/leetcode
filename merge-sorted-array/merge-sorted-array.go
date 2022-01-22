func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 1 {
		if len(nums2) != 0 {
			nums1[0] = nums2[0]
		}
	} else if len(nums2) > 0 {
		tmp := make([]int, m+n)
		vec1, vec2 := 0, 0
		for i := 0; i < m+n; i++ {
			if nums1[vec1] == 0 {
				if vec1 < m && nums2[vec2] >= 0 {
					tmp[i] = nums1[vec1]
					vec1++
				} else if vec2 < n {
					tmp[i] = nums2[vec2]
					vec2++
				}
			} else {
				if vec1 == m {
					tmp[i] = nums2[vec2]
					vec2++
				} else if vec2 == n {
					tmp[i] = nums1[vec1]
					vec1++
				} else if nums1[vec1] <= nums2[vec2] {
					tmp[i] = nums1[vec1]
					vec1++
				} else {
					tmp[i] = nums2[vec2]
					vec2++
				}
			}
		}
		copy(nums1[:], tmp[:])
	}
}