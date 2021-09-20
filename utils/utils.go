package utils

import "math"

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

func MaxF64(list []float64) float64 {
	max := math.Inf(-1)

	for _, val := range list {
		if val > max {
			max = val
		}
	}

	return max
}

func MinF64(list []float64) float64 {
	min := math.Inf(1)

	for _, val := range list {
		if val < min {
			min = val
		}
	}

	return min
}