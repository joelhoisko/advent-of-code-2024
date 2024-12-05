package main

import (
	"fmt"
	"regexp"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

func main() {
	lines, err := data.ReadXMAS()
	if err != nil {
		fmt.Println("whoops:", err)
		return
	}

	xCount := 0
	fCount := 0
	bCount := 0
	uCount := 0
	dCount := 0
	dlCount := 0
	drCount := 0
	ulCount := 0
	urCount := 0
	re := regexp.MustCompile(`X`)
	for lineIndex, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		indexes := []int{}
		// just get the first index, we don't care about the ending as it's just 1 char
		for _, row := range matches {
			indexes = append(indexes, row[0])
		}

		for _, index := range indexes {

			// check forward
			if index < len(line)-3 && line[index:index+4] == "XMAS" {
				xCount++
				fCount++
			}

			// check backwards
			if index > 2 && line[index-3:index+1] == "SAMX" {
				xCount++
				bCount++
			}
			// check straight down
			if lineIndex < len(lines)-3 {
				checkString := string(line[index]) + string(lines[lineIndex+1][index]) + string(lines[lineIndex+2][index]) + string(lines[lineIndex+3][index])
				if checkString == "XMAS" {
					xCount++
					dCount++
				}
			}
			// check straight up
			if lineIndex > 2 {
				checkString := string(line[index]) + string(lines[lineIndex-1][index]) + string(lines[lineIndex-2][index]) + string(lines[lineIndex-3][index])
				if checkString == "XMAS" {
					xCount++
					uCount++
				}
			}
			// check diagonal down left
			if index > 2 && lineIndex < len(lines)-3 {
				checkString := string(line[index]) + string(lines[lineIndex+1][index-1]) + string(lines[lineIndex+2][index-2]) + string(lines[lineIndex+3][index-3])
				if checkString == "XMAS" {
					xCount++
					dlCount++
				}
			}
			// check diagonal down right
			if index < len(lines)-3 && lineIndex < len(lines)-3 {
				checkString := string(line[index]) + string(lines[lineIndex+1][index+1]) + string(lines[lineIndex+2][index+2]) + string(lines[lineIndex+3][index+3])
				if checkString == "XMAS" {
					xCount++
					drCount++
				}
			}
			// check diagonal up left
			if index > 2 && lineIndex > 2 {
				checkString := string(line[index]) + string(lines[lineIndex-1][index-1]) + string(lines[lineIndex-2][index-2]) + string(lines[lineIndex-3][index-3])
				if checkString == "XMAS" {
					xCount++
					ulCount++
				}
			}
			// check diagonal up right
			if index < len(lines)-3 && lineIndex > 2 {
				checkString := string(line[index]) + string(lines[lineIndex-1][index+1]) + string(lines[lineIndex-2][index+2]) + string(lines[lineIndex-3][index+3])
				if checkString == "XMAS" {
					xCount++
					urCount++
				}
			}
		}
	}
	fmt.Println("XMAS count:", xCount)
	fmt.Println("Forward count:", fCount)
	fmt.Println("Backward count:", bCount)
	fmt.Println("Down count:", dCount)
	fmt.Println("Up count:", uCount)
	fmt.Println("Down left count:", dlCount)
	fmt.Println("Down right count:", drCount)
	fmt.Println("Up left count:", ulCount)
	fmt.Println("Up right count:", urCount)
}
