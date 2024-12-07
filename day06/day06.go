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

  var visitedPos, _ = getVisitedPositions(mapGrid, startPos)
	var part1Count int = len(visitedPos)
	var part2Count int = getObstacleInfiniteLoopCount(mapGrid, startPos, visitedPos)
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

func getObstacleInfiniteLoopCount(mapGrid [][]byte, startPos Position, visitedPos map[Position]bool) int {
	var obstacleCount int
  // Remove the starting position from the list of visited places we can test an obstacle at
  delete(visitedPos, Position{startPos.y, startPos.x})
  for pos := range visitedPos{
    testGrid := newGrid(mapGrid)
    testGrid[pos.y][pos.x] = '#'
    var _, isALoop = getVisitedPositions(testGrid, startPos)
    if isALoop {
      obstacleCount++
    }
  }
	return obstacleCount
}

func getVisitedPositions(mapGrid [][]byte, currentPos Position) (map[Position]bool, bool) {
	directions := map[byte]Position{
		'^': {-1, 0},
		'>': {0, 1},
		'V': {1, 0},
		'<': {0, -1},
	}
	turnOrder := []byte{'^', '>', 'V', '<'}
	currentDir := mapGrid[currentPos.y][currentPos.x]
	visitedPos := make(map[Position]bool)
	obstaclesHit := make(map[Position]byte)
  isALoop := false

	for {
		visitedPos[currentPos] = true
		nextPos := Position{currentPos.y + directions[currentDir].y, currentPos.x + directions[currentDir].x}
		if nextPos.y < 0 || nextPos.y >= len(mapGrid) || nextPos.x < 0 || nextPos.x >= len(mapGrid[0]) {
			break
		}
		if mapGrid[nextPos.y][nextPos.x] == '#' {
      if obstaclesHit[nextPos] == currentDir {
        isALoop = true
        return visitedPos, isALoop
      } else {
        obstaclesHit[nextPos] = currentDir
      }
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
	return visitedPos, isALoop
}

// Part 1 Notes
// ------------------------
// Map ^>V< to direction vectors
// Create map of visited coords
// Read map from txt file
// Find starting point (contains ^>V<)
// While not out of bounds
// Mark current location as visited
// Add to count
// Check in direction ahead for obstacle "#"
// If not, move forward to next location
// Else change direction to next in map

// Part 2 Notes
// ------------------------
// Only need to check locations where you initially walked
// Store location and direction of encounter with each obstacle
// Generate and test a new mapGrid with the new obstacle in place
// If you ever encounter an obstacle from the same direction, you must be stuck in a loop
// Therefore increment the loop counter and test the next location
