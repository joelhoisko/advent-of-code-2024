// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

// Main function
func main() {
	slices.Sort(data.LeftList)
	slices.Sort(data.RightList)

	var diffSum int
	for i, _ := range data.LeftList {
		difference := diff(data.LeftList[i], data.RightList[i])
		// fmt.Println("diff is: ", difference)
		diffSum += difference
	}
	fmt.Println("diff sum is: ", diffSum)
}

func findSmallest(list []int) (int, int) {
	smallest := list[0]
	smallestIndex := 0

	for i, num := range data.LeftList[1:] {
		if num < smallest {
			smallest = num
			smallestIndex = i
		}
	}

	return smallestIndex, smallest
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
