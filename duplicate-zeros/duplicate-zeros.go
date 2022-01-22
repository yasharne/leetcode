func duplicateZeros(arr []int) {
	var modified bool
	for i, s := range arr {
		if s == 0 && i+2 <= len(arr) {
			if modified {
				modified = false
				continue
			}
			copy(arr[i+2:], arr[i+1:])
			arr[i+1] = 0
			modified = true
		}
	}
}