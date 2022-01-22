func checkIfExist(arr []int) bool {
	for k := range arr {
		for i, j := range arr[k:] {
			if i != j && (float64(arr[k]) == float64(j)/2 || arr[k] == j*2) {
				return true
			}
		}
	}
	return false
}
