package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()

	fmt.Println("Part 1 Result:", partOne(input))
	fmt.Println("Part 2 Result:", partTwo(input))
}

func partOne(input []string) int {
	// All directions are valid for this part
	directions := []struct{ x, y int }{
		{x: 0, y: 1},   // Right
		{x: 1, y: 0},   // Down
		{x: 1, y: 1},   // Down Right
		{x: -1, y: 1},  // Up Right
		{x: 0, y: -1},  // Left
		{x: -1, y: 0},  // Up
		{x: -1, y: -1}, // Up Left
		{x: 1, y: -1},  // Down Left
	}
	return search(input, "XMAS", directions)
}

func partTwo(input []string) int {
	return 0
}

// Search for a word in a slice of strings considering all directions and forward/backwards.
// Returns the number of times the word was found.
func search(input []string, word string, directions []struct{ x, y int }) int {
	// First, find all of the coordinates of the first letter of the word.
	firstLetterCoordinates := findCoordinates(input, word[0])

	matches := 0

	// For each coordinate, check if the word is found in any direction.
	for _, coordinates := range firstLetterCoordinates {
		for _, direction := range directions {
			if checkWordInDirection(input, word, coordinates, direction) {
				matches++
			}
		}
	}

	return matches
}

// checkWordInDirection checks if the word is found in a specific direction starting from a coordinate.
func checkWordInDirection(input []string, word string, startCoord struct{ x, y int }, direction struct{ x, y int }) bool {
	x, y := startCoord.x, startCoord.y

	for i := 0; i < len(word); i++ {
		if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) || input[x][y] != word[i] {
			return false
		}
		// Move to the next character in the given direction
		x += direction.x
		y += direction.y
	}

	return true
}

// Find the coordinates of the given letter in the 2d slice.
func findCoordinates(input []string, letter byte) []struct{ x, y int } {
	coordinates := []struct{ x, y int }{}

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == letter {
				coordinates = append(coordinates, struct{ x, y int }{x: row, y: col})
			}
		}
	}

	return coordinates
}

// Read the input file and return the contents as a slice of strings.
func readInput() []string {
	lines := []string{}

	file, _ := os.Open("./input.txt")

	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		lines = append(lines, scanner.Text())
	}

	return lines
}
