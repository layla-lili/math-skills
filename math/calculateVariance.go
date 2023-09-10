package mathskills

// Function to calculate the variance
func CalculateVariance(numbers []float64) float64 {
	average := CalculateAverage(numbers)
	sumOfSquaredDifferences := 0.0
	for _, num := range numbers {
		diff := num - average
		sumOfSquaredDifferences += diff * diff
	}
	return sumOfSquaredDifferences / float64(len(numbers))
}
