package main

import (
	"fmt"
	"slices"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

func main() {
	instructions, pages, err := data.ReadPages()
	if err != nil {
		fmt.Println("whoops:", err)
		return
	}

	var correctPages [][]int
	for _, page := range pages {

		var relevantInstructions [][]int
		for _, instruction := range instructions {
			if slices.Contains(page, instruction[0]) && slices.Contains(page, instruction[1]) {
				relevantInstructions = append(relevantInstructions, instruction)
			}
		}

		isCorrect := true
		for numIndex, num := range page {
			for _, instruction := range relevantInstructions {
				if numIndex == 0 && num == instruction[1] {
					// it's so 🅱over
					isCorrect = false
					// fmt.Println("is over with num", num, "due to instruction", instruction)
					break
				} else if numIndex > 0 {
					// aka 75,95 -> 95|75
					if instruction[0] == num && page[numIndex-1] == instruction[1] {
						// it's so 🅱over
						isCorrect = false
						// fmt.Println("is over with num", page[numIndex-1], "and num", num, "due to instruction", instruction)
						break
					}
				}
			}
		}
		if isCorrect {
			correctPages = append(correctPages, page)
		}
	}
	sum := 0
	fmt.Println("Correct pages:", len(correctPages))
	for _, page := range correctPages {
		sum += page[len(page)/2]
	}

	fmt.Println("Sum of middle pages:", sum)
}
