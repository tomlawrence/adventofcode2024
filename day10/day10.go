package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row, col int
}

func main() {
	file, err := os.Open("day10/day10.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mapGrid := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var lineNums []int
		for i := 0; i < len(line); i++ {
			num, err := strconv.Atoi(line[i])
			if err == nil {
				lineNums = append(lineNums, num)
			}
		}
		mapGrid = append(mapGrid, lineNums)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var part1, part2 int = sumTrailCounts(mapGrid)
	fmt.Println("Part 1 - Sum of scores (number of 9s that can be reached):", part1)
	fmt.Println("Part 2 - Sum of ratings (number of distinct trails to 9): ", part2)
}

func sumTrailCounts(mapGrid [][]int) (int, int) {
	var part1, part2 int
	for row := range mapGrid {
		for col := range mapGrid[row] {
			if mapGrid[row][col] == 0 {
				var visitedEnds = make(map[Position]bool)
				part1 += findTrails(mapGrid, row, col, 1, visitedEnds)
				part2 += len(visitedEnds)
			}
		}
	}
	return part1, part2
}

func findTrails(mapGrid [][]int, row, col, nextNum int, visitedEnds map[Position]bool) int {
	var trailCount int
	var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	if nextNum == 10 {
		visitedEnds[Position{row, col}] = true
		return 1
	}
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < len(mapGrid) && newCol >= 0 && newCol < len(mapGrid[0]) {
			if mapGrid[newRow][newCol] == nextNum {
				trailCount += findTrails(mapGrid, newRow, newCol, nextNum+1, visitedEnds)
			}
		}
	}
	return trailCount
}
