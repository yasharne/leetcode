func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var flip = false
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == arr[i] || arr[i] == arr[i+1] {
			return false
		}
		if arr[i-1] > arr[i] && arr[i] < arr[i+1] {
			return false
		}
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			flip = true
		}
	}
	if flip {
		return true
	}
	return false
}