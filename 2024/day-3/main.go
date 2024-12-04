package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Println("Part 1 Result:", partOne(input))
	fmt.Println("Part 2 Result:", partTwo(input))
}

func partOne(input string) int {
	operations := findOperations(input)

	sum := 0
	for _, operation := range operations {
		a, b := findFactors(operation)
		sum += a * b
	}

	return sum
}

func partTwo(input string) int {
	operations := filterOperations(findOperationsWithConditional(input))

	sum := 0
	for _, operation := range operations {
		a, b := findFactors(operation)
		sum += a * b
	}
	return sum
}

func findOperations(input string) []string {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	return pattern.FindAllString(input, -1)
}

func findOperationsWithConditional(input string) []string {
	// This pattern will match mul, don't, and do operations in the order they appear in the string.
	pattern := regexp.MustCompile(`(mul\(\d+,\d+\)|don\'t\(\)|do\(\))`)

	return pattern.FindAllString(input, -1)
}

func filterOperations(operations []string) []string {
	filtered := []string{}

	skip := false

	for _, operation := range operations {
		switch operation {
		case "don't()":
			skip = true
		case "do()":
			skip = false
		default:
			if !skip {
				filtered = append(filtered, operation)
			}
		}
	}

	return filtered
}

func findFactors(operation string) (int, int) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindStringSubmatch(operation)

	a, _ := strconv.Atoi(matches[1])
	b, _ := strconv.Atoi(matches[2])

	return a, b
}

// Read the input file and return the contents as a string.
func readInput() string {
	bytes, _ := os.ReadFile("./input.txt")

	return string(bytes)
}
