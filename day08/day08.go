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
	file, err := os.Open("day08/day08.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mapGrid := [][]byte{}
	antennas := map[byte][]Position{}
	var antennaCount int
	var y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		mapGrid = append(mapGrid, row)
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Position{y, x})
				antennaCount++
			}
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Antenna count:", antennaCount)
	for antenna, coords := range antennas {
		fmt.Println(string(antenna), coords)
	}
	part1Count := getAntinodeCount(mapGrid, antennas)
	fmt.Println("Part 1 - Number of antinodes:", part1Count)
}

func getAntinodeCount(mapGrid [][]byte, antennas map[byte][]Position) int {
	var antinodeCount int
	antinodes := make(map[Position]bool)

	for antenna, coords := range antennas {
		fmt.Println(string(antenna))
		for i, pos1 := range coords[:len(coords)-1] {
			for _, pos2 := range coords[i+1:] {
				delta := Position{pos2.y - pos1.y, pos2.x - pos1.x}
				dPos1 := Position{pos1.y - delta.y, pos1.x - delta.x}
				dPos2 := Position{pos2.y + delta.y, pos2.x + delta.x}
				if isInBoundaries(dPos1, mapGrid) {
					if !antinodes[dPos1] {
						fmt.Println("Antinode of", string(antenna), "at", dPos1, "with delta", delta)
						antinodes[dPos1] = true
						antinodeCount++
					}
				}
				if isInBoundaries(dPos2, mapGrid) {
					if !antinodes[dPos2] {
						fmt.Println("Antinode of", string(antenna), "at", dPos2, "with delta", delta)
						antinodes[dPos2] = true
						antinodeCount++
					}
				}
			}
		}
	}
	return antinodeCount
}

func isInBoundaries(pos Position, mapGrid [][]byte) bool {
	return pos.y >= 0 && pos.y < len(mapGrid) && pos.x >= 0 && pos.x < len(mapGrid[0])
}
