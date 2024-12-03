package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := readInput()

	partOne(reports)
	partTwo(reports)
}

func partOne(reports [][]int) {
	validCount := 0

	for _, report := range reports {
		if isReportValid(report) {
			validCount++
		}
	}

	fmt.Println("Part 1 Result:", validCount)
}

func partTwo(reports [][]int) {
	validCount := 0

	for _, report := range reports {
		for _, combination := range getCombinations(report) {
			if isReportValid(combination) {
				validCount++
				// Only count once if any combination is valid
				break
			}
		}
	}

	fmt.Println("Part 2 Result:", validCount)
}

// Generate all combinations by removing one element at a time from a given slice.
// For example, given [1, 2, 3], the combinations would be [[2, 3], [1, 3], [1, 2]].
func getCombinations(arr []int) [][]int {
	combinations := make([][]int, len(arr))

	for i := range arr {
		combinations[i] = append([]int{}, arr[:i]...)
		combinations[i] = append(combinations[i], arr[i+1:]...)
	}

	return combinations
}

// Check if a report is valid by ensuring that the difference between each neighbor is at most 3
// and the report is either ascending or descending.
func isReportValid(report []int) bool {
	ascending, descending := true, true

	for i := 0; i < len(report)-1; i++ {
		diff := absDiff(report[i], report[i+1])
		if diff > 3 || diff == 0 {
			return false
		}
		if report[i] < report[i+1] {
			descending = false
		} else if report[i] > report[i+1] {
			ascending = false
		}
	}

	return ascending || descending
}

// Calculate the absolute difference between two integers.
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// Read the input file and parse the reports into a 2D slice of integers.
func readInput() [][]int {
	var reports [][]int

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		parsedReport := make([]int, len(columns))
		for i, column := range columns {
			num, _ := strconv.Atoi(column)
			parsedReport[i] = num
		}

		reports = append(reports, parsedReport)
	}

	return reports
}
