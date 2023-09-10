package main

import (
	"fmt"
	"log"
	"math"
	mathskills "mathskills/math"
	"os"
)

func main() {
	// Specify the file path
	filePath := "data.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the numbers from the file
	numbers, err := mathskills.ReadNumbersFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate average
	average :=round(mathskills.CalculateAverage(numbers))
	fmt.Println("Average: ", average)

	// Calculate median
	median := round(mathskills.CalculateMedian(numbers))
	fmt.Println("Median: ", median)
	// Calculate variance
	variance := round(mathskills.CalculateVariance(numbers))
	fmt.Println("Variance: ", variance)

	// Calculate standard deviation
	standardDeviation := round(mathskills.CalculateStandardDeviation(numbers))
	fmt.Println("StandardDeviation: ", standardDeviation)
}

// Function to round a float64 value to the nearest integer
func round(value float64) int {
	return int(math.Round(value))
}