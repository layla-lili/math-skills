package mathskills

import "math"

// Function to calculate the standard deviation
func CalculateStandardDeviation(numbers []float64) float64 {
	variance := CalculateVariance(numbers)
	return math.Sqrt(variance)
}