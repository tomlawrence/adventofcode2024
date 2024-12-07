package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	y, x int
}

func main() {
	file, err := os.Open("day06/day06.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mapGrid := [][]byte{}
	var startFound bool
	var startPos Position
	var y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		mapGrid = append(mapGrid, row)
		if !startFound {
			for x, char := range row {
				if char == '^' || char == '>' || char == 'V' || char == '<' {
					startPos = Position{y, x}
					startFound = true
					break
				}
			}
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if !startFound {
		fmt.Println("Unable to find starting position in input")
		return
	}

	var part1Count int = getVisitedPositionsCount(mapGrid, startPos)
	var part2Count int = getObstacleInfiniteLoopCount(mapGrid, startPos)
	fmt.Println("Part 1 - Number of locations visited by guard:", part1Count)
	fmt.Println("Part 2 - Number of infinite loop obstacle locations:", part2Count)
}

func newGrid(mapGrid [][]byte) [][]byte {
	newGrid := make([][]byte, len(mapGrid))
	for i := range mapGrid {
		newGrid[i] = make([]byte, len(mapGrid[i]))
		copy(newGrid[i], mapGrid[i])
	}
	return newGrid
}

func getObstacleInfiniteLoopCount(mapGrid [][]byte, startPos Position) int {
	var obstacleCount int

	for y, row := range mapGrid {
		for x, col := range row {
			if col == '^' || col == '>' || col == 'V' || col == '<' || col == '#' {
			} else {
				testGrid := newGrid(mapGrid)
				testGrid[y][x] = '#'
				if getVisitedPositionsCount(testGrid, startPos) == 30000 {
					obstacleCount++
				}
			}
			// if col != '#' || col != '^' || col != '>' || col != 'V' || col != '<' {
			// 	testGrid := newGrid(mapGrid)
			// 	testGrid[y][x] = '#'
			// 	if getVisitedPositionsCount(testGrid, startPos) == 25000 {
			//        fmt.Println("Obstacle at ", y, x)
			// 		obstacleCount++
			// 	}
			// }
		}
	}

	return obstacleCount
}

func getVisitedPositionsCount(mapGrid [][]byte, currentPos Position) int {
	directions := map[byte]Position{
		'^': {-1, 0},
		'>': {0, 1},
		'V': {1, 0},
		'<': {0, -1},
	}
	turnOrder := []byte{'^', '>', 'V', '<'}
	stepCount := 0

	currentDir := mapGrid[currentPos.y][currentPos.x]
	visitedPos := make(map[Position]bool)

	for {
		stepCount++
		if stepCount == 30000 {
			return stepCount
		}
		visitedPos[currentPos] = true
		nextPos := Position{currentPos.y + directions[currentDir].y, currentPos.x + directions[currentDir].x}
		if nextPos.y < 0 || nextPos.y >= len(mapGrid) || nextPos.x < 0 || nextPos.x >= len(mapGrid[0]) {
			break
		}
		// var confirmedClear bool
		if mapGrid[nextPos.y][nextPos.x] == '#' {
			for i, dir := range turnOrder {
				if dir == currentDir {
					currentDir = turnOrder[(i+1)%4]
					break
				}
			}
			continue
		}
		currentPos = nextPos
	}
	return len(visitedPos)
}

/* Map ^>V< to direction vectors
create map of visited coords
read map from txt file
find starting point (contains ^>V<)
while not out of bounds
mark current location as visited
add to count
check in direction ahead for obstacle "#"
if not, move forward to next location
else change direction to next in map
*/
