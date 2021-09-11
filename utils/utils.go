package utils

// Sums all of the numbers in an integer array
func Sum(nums []int) int {
	total := 0
	for _, num := range(nums) {
		total += num
	}

	return total
}

// Check if target is between lower and upper for Float64
func BetweenFloat64(lower float64, upper float64, target float64) bool {
	if target < lower || target > upper {
		return false
	}

	return true
}