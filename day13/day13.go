package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

type Claw struct {
	A, B, Prize Position
}

func main() {
	file, err := os.ReadFile("day13/day13.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	clawMachines := []Claw{}
	clawString := strings.Split(string(file), "\n\n")
	for _, claw := range clawString {
		clawRegex := regexp.MustCompile(`(\d+)`)
		coords := clawRegex.FindAllStringSubmatch(claw, 6)
		aX, err1 := strconv.Atoi(coords[0][0])
		aY, err2 := strconv.Atoi(coords[1][0])
		bX, err3 := strconv.Atoi(coords[2][0])
		bY, err4 := strconv.Atoi(coords[3][0])
		pX, err5 := strconv.Atoi(coords[4][0])
		pY, err6 := strconv.Atoi(coords[5][0])
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			fmt.Println("Error parsing claw machine coordinates to ints")
			return
		}
		clawMachines = append(clawMachines, Claw{Position{aX, aY}, Position{bX, bY}, Position{pX, pY}})
	}
	// for _, claw := range clawMachines {
	// 	fmt.Printf("Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d\n\n", claw.A.X, claw.A.Y, claw.B.X, claw.B.Y, claw.Prize.X, claw.Prize.Y)
	// }
	fmt.Println("Part 1 - Minimum tokens required: ", minimumTokensRequired(clawMachines, false))
	fmt.Println("Part 2 - Minimum tokens required: ", minimumTokensRequired(clawMachines, true))
}

func minimumTokensRequired(clawMachines []Claw, part2 bool) int {
	var minTokens int = 0
	for _, claw := range clawMachines {
		a, b := minimumButtonPresses(claw, part2)
		minTokens += (a * 3) + b
	}
	return minTokens
}

func minimumButtonPresses(claw Claw, part2 bool) (int, int) {
	if part2 {
		claw.Prize.X += 10000000000000
		claw.Prize.Y += 10000000000000
	}
	// Cramer's rule to solve simultaneous linear equations
	// Here's a plot example: https://www.wolframalpha.com/input?i=94a%2B22b%3D8400%2C+34a%2B67b%3D5400 (copy/paste URL)
	det := claw.A.X*claw.B.Y - claw.A.Y*claw.B.X
	detA := claw.Prize.X*claw.B.Y - claw.Prize.Y*claw.B.X
	detB := claw.Prize.Y*claw.A.X - claw.Prize.X*claw.A.Y

	if detA%det != 0 && detB%det != 0 {
		return 0, 0
	}
	a := detA / det
	b := detB / det
	// fmt.Println("det:", det, "\ta:", a, "\tb:", b)
	return a, b
}
