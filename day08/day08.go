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
	part1Count := getAntinodeCount(mapGrid, antennas, false)
	part2Count := getAntinodeCount(mapGrid, antennas, true)
	fmt.Println("Part 1 - Number of antinodes:", part1Count)
	fmt.Println("Part 2 - Number of antinodes:", part2Count)
}

func getAntinodeCount(mapGrid [][]byte, antennas map[byte][]Position, part2 bool) int {
	antinodes := make(map[Position]bool)

	for antenna, coords := range antennas {
		for i, pos1 := range coords[:len(coords)-1] {
			for _, pos2 := range coords[i+1:] {
				if part2 {
					antinodes[pos2] = true
					antinodes[pos1] = true
				}

				delta := Position{pos2.y - pos1.y, pos2.x - pos1.x}
				dPos1 := Position{pos1.y - delta.y, pos1.x - delta.x}
				dPos2 := Position{pos2.y + delta.y, pos2.x + delta.x}

				for {
					if isInBoundaries(dPos1, mapGrid) {
						if !antinodes[dPos1] {
							fmt.Println("Antinode of", string(antenna), "at", dPos1, "with delta", delta)
							antinodes[dPos1] = true
						}
						if part2 {
							dPos1 = Position{dPos1.y - delta.y, dPos1.x - delta.x}
						} else {
							break
						}
					} else {
						fmt.Println("Out of bounds: Antinode of", string(antenna), "at", dPos1, "with delta", delta)
						break
					}
				}

				for {
					if isInBoundaries(dPos2, mapGrid) {
						if !antinodes[dPos2] {
							fmt.Println("Antinode of", string(antenna), "at", dPos2, "with delta", delta)
							antinodes[dPos2] = true
						}
						if part2 {
							dPos2 = Position{dPos2.y + delta.y, dPos2.x + delta.x}
						} else {
							break
						}
					} else {
						fmt.Println("Out of bounds: Antinode of", string(antenna), "at", dPos2, "with delta", delta)
						break
					}
				}

			}
		}
	}
	return len(antinodes)
}

func isInBoundaries(pos Position, mapGrid [][]byte) bool {
	return pos.y >= 0 && pos.y < len(mapGrid) && pos.x >= 0 && pos.x < len(mapGrid[0])
}
