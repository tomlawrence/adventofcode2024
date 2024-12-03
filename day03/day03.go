package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("day03/day03.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	var instructions string = string(file)
  fmt.Println("Part 1 Multiplication Results:", multiply(instructions, false))
  fmt.Println("Part 2 Multiplication Results:", multiply(instructions, true))
}

func multiply(instructions string, conditionalsEnabled bool) int {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
  if conditionalsEnabled {
		mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	}
	var results int
	var enabled bool = true
	cleanInstructions := mulRegex.FindAllStringSubmatch(instructions, -1)
	for _, instruction := range cleanInstructions {
		switch instruction[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				num1, err1 := strconv.Atoi(instruction[1])
				num2, err2 := strconv.Atoi(instruction[2])
				if err1 == nil && err2 == nil {
					results += num1 * num2
				} else {
          fmt.Println("Error parsing multiply insruction to ints", num1, num2)
				}
			}
		}
	}
  return results
}
