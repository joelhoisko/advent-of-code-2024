package data

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPages() ([][]int, [][]int, error) {
	file, err := os.Open("data/day5.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var instructions [][]int
	var pages [][]int

	instructionMode := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into fields (e.g., by spaces)
		line := scanner.Text()
		if line == "" {
			instructionMode = false
			continue
		}
		if instructionMode {
			fields := strings.Split(line, "|")

			// Convert the fields to integers
			var row []int
			for _, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					return nil, nil, fmt.Errorf("error converting field to integer for instruction	: %w", err)
				}
				row = append(row, num)
			}

			// Append the row to the matrix
			instructions = append(instructions, row)
		} else {
			fields := strings.Split(line, ",")

			// Convert the fields to integers
			var row []int
			for _, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					return nil, nil, fmt.Errorf("error converting field to integer: %w", err)
				}
				row = append(row, num)
			}

			// Append the row to the matrix
			pages = append(pages, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return instructions, pages, nil
}
