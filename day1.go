package main

import (
	"fmt"
	"slices"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

func main() {
	slices.Sort(data.LeftList)
	slices.Sort(data.RightList)

	var diffSum int
	for i, _ := range data.LeftList {
		difference := diff(data.LeftList[i], data.RightList[i])
		diffSum += difference
	}
	fmt.Println("diff sum is: ", diffSum)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
