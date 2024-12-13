package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	y, x int
}

type Region struct {
	Type      byte
	Plants    []Position
	Perimeter int
	Edges     int
}

func main() {
	file, err := os.Open("day12/day12.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	garden := [][]byte{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		garden = append(garden, append([]byte{}, scanner.Bytes()...))
	}
	calculateTotalPrice(garden)
}

func findRegion(start Position, garden [][]byte, visited map[Position]bool) Region {
	region := Region{
		Type:   garden[start.y][start.x],
		Plants: []Position{start},
	}
	directions := []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// BFS search for adjacent plants
	visited[start] = true
	queue := []Position{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			y := current.y + dir.y
			x := current.x + dir.x
			// Count perimeter if pos is outside grid or contains different plant type
			if y < 0 || y >= len(garden) || x < 0 || x >= len(garden[0]) || garden[y][x] != region.Type {
				region.Perimeter++
				continue
			}
			// If it's the same plant type and not visited, add to queue
			if !visited[Position{y, x}] {
				visited[Position{y, x}] = true
				region.Plants = append(region.Plants, Position{y, x})
				queue = append(queue, Position{y, x})
			}
		}
	}
	return region
}

func findAllRegions(garden [][]byte) []Region {
	regions := []Region{}
	visited := make(map[Position]bool)

	for y, row := range garden {
		for x, _ := range row {
			pos := Position{y, x}
			if !visited[pos] {
				regions = append(regions, findRegion(pos, garden, visited))
			}
		}
	}
	return regions
}

func calculateTotalPrice(garden [][]byte) int {
	totalPrice := 0
	for _, region := range findAllRegions(garden) {
		totalPrice += len(region.Plants) * region.Perimeter
		//fmt.Println(string(region.Type), len(region.Plants), "*", region.Perimeter)
	}
	fmt.Println("Part 1 - Total Price:", totalPrice)
	return totalPrice
}
