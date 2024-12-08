package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	target  int
	numbers []int
}

func main() {
	file, err := os.Open("day07/day07.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sums := []Case{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		target, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error converting target number: %v\n", err)
			continue
		}
		numsString := strings.Split(line[1], " ")
		var nums []int
		for i := 0; i < len(numsString); i++ {
			num, err := strconv.Atoi(numsString[i])
			if err != nil {
				fmt.Printf("Error converting number: %v\n", err)
				continue
			}
			nums = append(nums, num)
		}
		sums = append(sums, Case{target: target, numbers: nums})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var part1Count, part1SumsTotal int = getCalibrationCountAndSum(sums, false)
	var part2Count, part2SumsTotal int = getCalibrationCountAndSum(sums, true)
	fmt.Printf("Part 1 - Valid target count: %v, Sum of valid targets: %v\n", part1Count, part1SumsTotal)
	fmt.Printf("Part 2 - Valid target count: %v, Sum of valid targets: %v\n", part2Count, part2SumsTotal)
}

func getCalibrationCountAndSum(sums []Case, useConcat bool) (int, int) {
	var count int
	var targetsSumTotal int
	for _, test := range sums {
		found, _ := findSequence(test.numbers, test.target, useConcat) // Unused return val is operations for debugging purposes
		// fmt.Printf("Numbers: %v, Target: %d\n", test.numbers, test.target)
		if found {
			// fmt.Printf("Solution found: %s\n", formatSolution(test.numbers, operations))
			targetsSumTotal += test.target
			count++
		} else {
			// fmt.Printf("No solution found\n")
		}
	}
	return count, targetsSumTotal
}

func findSequence(numbers []int, target int, useConcat bool) (bool, []string) {
	operations := make([]string, len(numbers)-1)
	var solve func(index int, currentTotal int) bool
	solve = func(index int, currentTotal int) bool {
		if index == len(numbers) {
			return currentTotal == target
		}
		if index == 0 {
			return solve(index+1, numbers[0])
		}
		operations[index-1] = "+"
		if solve(index+1, currentTotal+numbers[index]) {
			return true
		}
		operations[index-1] = "*"
		if solve(index+1, currentTotal*numbers[index]) {
			return true
		}
		if useConcat {
			operations[index-1] = "||"
			if solve(index+1, concatenateNumbers(currentTotal, numbers[index])) {
				return true
			}
		}
		return false
	}
	found := solve(0, 0)
	if !found {
		return false, nil
	}
	return true, operations
}

func formatSolution(numbers []int, operations []string) string {
	if len(operations) == 0 {
		return fmt.Sprintf("%d", numbers[0])
	}
	result := fmt.Sprintf("%d", numbers[0])
	for i, op := range operations {
		result += fmt.Sprintf(" %s %d", op, numbers[i+1])
	}
	return result
}

func concatenateNumbers(a, b int) int {
	// Convert b to string to get its length
	bStr := strconv.Itoa(b)
	// Multiply a by 10^(length of b) and add b
	multiplier := 1
	for i := 0; i < len(bStr); i++ {
		multiplier *= 10
	}
	return a*multiplier + b
}
