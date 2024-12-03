package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	left, right := readInputColumns()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += absDiff(left[i], right[i])
	}

	fmt.Println("Part 1 Result:", sum)
}

func partTwo() {
	left, right := readInputColumns()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		occurrences := countOccurences(left[i], right)
		similarityScore := occurrences * left[i]

		sum += similarityScore
	}

	fmt.Println("Part 2 Result:", sum)
}

func countOccurences(value int, arr []int) int {
	count := 0
	for _, v := range arr {
		if v == value {
			count++
		}
	}

	return count
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func readInputColumns() ([]int, []int) {
	left := []int{}
	right := []int{}

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		columns := strings.Fields(line)

		leftInt, _ := strconv.Atoi(columns[0])
		rightInt, _ := strconv.Atoi(columns[1])

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	return left, right
}
