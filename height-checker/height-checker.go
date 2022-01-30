func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	_ = copy(sorted, heights)
	sort.Ints(sorted)
	j := 0
	for i := 0; i < len(heights); i++ {
		if sorted[i] != heights[i] {
			j++
		}
	}
	return j
}