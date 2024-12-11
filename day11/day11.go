package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Stone struct {
	Data int
	Size int
}

func main() {
	file, err := os.ReadFile("day11/day11.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	stonesString := strings.Split(string(file), " ")
	var stones []int
	for _, char := range stonesString {
		stone, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
			continue
		}
		stones = append(stones, stone)
	}
	fmt.Println(stones)
	fmt.Println("Part 1 - Number of stones:", part1(stones, 25))
}

func part1(stones []int, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func blink(stones []int) []int {
	stonesNew := append([]int{}, stones...)
	var j int
	for _, stone := range stones {
		switch {
		case stone == 0:
			stonesNew[j] = 1
		case isEven(stone):
			left, right := splitNumber(stone)
			stonesNew[j] = right
			stonesNew = slices.Insert(stonesNew, j, left)
			j++
		default:
			stonesNew[j] *= 2024
		}
		j++
	}
	//fmt.Println(stonesNew)
	return stonesNew
}

func isEven(n int) bool {
	digits := 1
	for n >= 10 {
		n /= 10
		digits++
	}
	return digits%2 == 0
}

func splitNumber(n int) (left, right int) {
	digits := 1
	for temp := n; temp >= 10; temp /= 10 {
		digits++
	}
	midpoint := (digits + 1) / 2
	divisor := int(math.Pow10(digits - midpoint))

	left = n / divisor
	right = n % divisor
	return left, right
}
