package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day04/day04.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chars := []byte(scanner.Text())
		grid = append(grid, chars)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	part1(grid)
	part2(grid)
}

func part1(grid [][]byte) {
	word := []byte("XMAS")
	fmt.Println("Part 1:", string(word), "found", findWord(word, grid, false), "times.")
}

func part2(grid [][]byte) {
	word := []byte("MAS")
	fmt.Println("Part 2: X-MAS found", findWord(word, grid, true), "times.")
}

func findWord(word []byte, grid [][]byte, part2 bool) int {
	var char byte = word[0]
	if part2 {
		char = word[1]
	}
	var count int
	var rows, cols int = len(grid), len(grid[0])
	var directions = [8][2]int{
		{-1, 0},  // Up
		{-1, 1},  // Up-Right
		{0, 1},   // Right
		{1, 1},   // Down-Right
		{1, 0},   // Down
		{1, -1},  // Down-Left
		{0, -1},  // Left
		{-1, -1}, // Left-Up
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == char {
				if !part2 {
					for _, direction := range directions {
						if checkDirection(direction, y, x, word, grid) {
							count++
						}
					}
				} else if checkCross(y, x, word, grid) {
					count++
				}
			}
		}
	}
	return count
}

func checkDirection(direction [2]int, startY int, startX int, word []byte, grid [][]byte) bool {
	var rows, cols int = len(grid), len(grid[0])

	// Multiply vector by word length to get end coords
	var endY int = startY + (direction[0] * (len(word) - 1))
	var endX int = startX + (direction[1] * (len(word) - 1))

	// Check coords don't exceed grid boundaries
	if endY >= rows || endY < 0 || endX >= cols || endX < 0 {
		return false
	}

	// Check each char in a given direction
	for i := 0; i < len(word); i++ {
		y := startY + (direction[0] * i)
		x := startX + (direction[1] * i)
		if grid[y][x] == word[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkCross(y int, x int, word []byte, grid [][]byte) bool {
	var rows, cols int = len(grid), len(grid[0])

	// Check surrounding coords don't exceed grid boundaries
	if y+1 >= rows || y-1 < 0 || x+1 >= cols || x-1 < 0 {
		return false
	}

	// Check diagonals in both directions for match
	if ((grid[y-1][x-1] == word[0] && grid[y+1][x+1] == word[2]) ||
		(grid[y-1][x-1] == word[2] && grid[y+1][x+1] == word[0])) &&
		((grid[y-1][x+1] == word[0] && grid[y+1][x-1] == word[2]) ||
			(grid[y-1][x+1] == word[2] && grid[y+1][x-1] == word[0])) {
		return true
	} else {
		return false
	}
}
