package utils

func InSlice(slice []string, x string) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}
	return false
}

func InSliceInt(slice []int, x int) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}
	return false
}
