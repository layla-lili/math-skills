package mathskills

import "sort"

// Function to calculate the median
func CalculateMedian(numbers []float64) float64 {
	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	// Calculate the median
	length := len(numbers)
	if length%2 == 0 {
		// If the length is even, average the middle two numbers
		middle := length / 2
		return (numbers[middle-1] + numbers[middle]) / 2
	} else {
		// If the length is odd, return the middle number
		middle := length / 2
		return numbers[middle]
	}
}