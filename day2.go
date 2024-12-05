package main

import (
	"fmt"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

func main() {
	matrix, err := data.ReadMatrix()
	if err != nil {
		fmt.Println("whoops:", err)
		return
	}

	safeLevels := 0
	for _, row := range matrix {
		isSafe := true
		isPositive := row[0]-row[1] > 0
		for i := 1; i < len(row); i++ {
			diff := row[i-1] - row[i]
			if diff > 0 != isPositive {
				fmt.Println("direction is fugged")
				isSafe = false
				break
			}
			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				fmt.Println("diff fugged")
				isSafe = false
				break
			}
		}
		if isSafe {
			fmt.Println("WE SAFE")
			safeLevels += 1
		}
	}

	fmt.Println("safe levels:", safeLevels)
}
