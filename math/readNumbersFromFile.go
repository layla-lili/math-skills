package mathskills

import (
	"bufio"
	"os"
	"strconv"
)

// Function to read numbers from a file
func ReadNumbersFromFile(file *os.File) ([]float64, error) {
	var numbers []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}