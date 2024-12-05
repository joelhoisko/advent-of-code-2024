package data

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadMatrix() ([][]int, error) {
	file, err := os.Open("data/day2.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var matrix [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into fields (e.g., by spaces)
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert the fields to integers
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("error converting field to integer: %w", err)
			}
			row = append(row, num)
		}

		// Append the row to the matrix
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return matrix, nil
}
