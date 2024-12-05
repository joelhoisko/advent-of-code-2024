package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joelhoisko/advent-of-code-2024/data"
)

func main() {
	lines, err := data.ReadFile()
	if err != nil {
		fmt.Println("whoops:", err)
		return
	}

	sum := 0
	for _, line := range lines {
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		for _, statement := range re.FindAllString(line, -1) {
			statement = statement[4 : len(statement)-1]
			strings := strings.Split(statement, ",")
			first, _ := strconv.Atoi(strings[0])
			second, _ := strconv.Atoi(strings[1])
			sum += first * second
		}
	}
	fmt.Println("sum:", sum)
}
