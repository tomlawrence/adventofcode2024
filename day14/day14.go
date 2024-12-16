package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	pX, pY int // position
	vX, vY int // velocity
}

func main() {
	file, err := os.Open("day14/day14.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	robots := []Robot{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		pString := strings.Split(strings.TrimPrefix(line[0], "p="), ",")
		vString := strings.Split(strings.TrimPrefix(line[1], "v="), ",")
		pX, err1 := strconv.Atoi(pString[0])
		pY, err2 := strconv.Atoi(pString[1])
		vX, err3 := strconv.Atoi(vString[0])
		vY, err4 := strconv.Atoi(vString[1])
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error converting robot value:", err1, err2, err3, err4)
			return
		}
		robots = append(robots, Robot{pX, pY, vX, vY})
	}
	var bathroom [][]int = make([][]int, 103)
	for i := range bathroom {
		bathroom[i] = make([]int, 101)
	}
	part1, part2 := parts1And2(robots, bathroom)
	fmt.Println("Part 1: Safety Factor (100 seconds):", part1)
	fmt.Println("Part 2: Number of seconds to display Christmas tree:", part2)
}

func parts1And2(robots []Robot, bathroom [][]int) (int, int) {
	var safetyFactor int
	var part2Seconds int
	for i := 0; i < 10000; i++ {
		robots = moveRobots(robots, bathroom)
		if i == 99 {
			q1, q2, q3, q4 := quadrantsRobotCount(robots, bathroom)
			safetyFactor = q1 * q2 * q3 * q4
		}
		//fmt.Println(i)
		//printBathroom(robots, bathroom)
		if hasChristmasTree(robots, bathroom) {
			part2Seconds = i + 1
			break
		}
	}
	return safetyFactor, part2Seconds
}

func moveRobots(robots []Robot, bathroom [][]int) []Robot {
	for i := range robots {
		newY := (robots[i].pY + (robots[i].vY)) % len(bathroom)
		newX := (robots[i].pX + (robots[i].vX)) % len(bathroom[0])
		if newY < 0 {
			newY = len(bathroom) + newY
		}
		if newX < 0 {
			newX = len(bathroom[0]) + newX
		}
		robots[i].pY = newY
		robots[i].pX = newX
	}
	return robots
}

func quadrantsRobotCount(robots []Robot, bathroom [][]int) (int, int, int, int) {
	var q1, q2, q3, q4 int
	var midY, midX int = (len(bathroom) / 2), (len(bathroom[0]) / 2)
	for _, robot := range robots {
		switch {
		case robot.pX < midX && robot.pY < midY:
			q1++
		case robot.pX > midX && robot.pY < midY:
			q2++
		case robot.pX < midX && robot.pY > midY:
			q3++
		case robot.pX > midX && robot.pY > midY:
			q4++
		}
	}
	return q1, q2, q3, q4
}

func hasChristmasTree(robots []Robot, bathroom [][]int) bool {
	for j := range bathroom {
		for i := range bathroom[j] {
			bathroom[j][i] = 0
		}
	}
	for _, robot := range robots {
		bathroom[robot.pY][robot.pX]++
	}
	count := 0
	for _, row := range bathroom {
		for _, val := range row {
			if val == 1 {
				count++
				if count == 10 { // Check if row contains a bunch of 1s in a row
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func printBathroom(robots []Robot, bathroom [][]int) {
	for j := range bathroom {
		for i := range bathroom[j] {
			bathroom[j][i] = 0
		}
	}
	for _, robot := range robots {
		bathroom[robot.pY][robot.pX]++
	}
	for _, row := range bathroom {
		for _, col := range row {
			if col != 0 {
				fmt.Printf("%d", col)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
