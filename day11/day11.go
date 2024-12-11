package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day11/day11.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	stonesString := strings.Split(string(file), " ")
	stones := make(map[int]int)

	for _, char := range stonesString {
		number, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
			continue
		}
		stones[number]++
	}
	fmt.Println("Part 1 - Number of stones:", blink(stones, 25))
	fmt.Println("Part 2 - Number of stones:", blink(stones, 75))
}

func blink(stones map[int]int, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = arrange(stones)
		//fmt.Printf("Loop %v: %v\n", i+1, countStones(stones))
	}
	return countStones(stones)
}

func arrange(stones map[int]int) map[int]int {
	stonesNew := make(map[int]int)
	for num, qty := range stones {
		stonesNew[num] = qty
	}
	for num, qty := range stones {
		switch {
		case num == 0:
			stonesNew[1] += qty
			stonesNew[0] -= qty
		case isEven(num):
			left, right := splitNumber(num)
			stonesNew[right] += qty
			stonesNew[left] += qty
			stonesNew[num] -= qty
		default:
			stonesNew[num] -= qty
			num *= 2024
			stonesNew[num] += qty
		}
	}
	return stonesNew
}

func countStones(stones map[int]int) int {
	var count int
	for _, qty := range stones {
		count += qty
	}
	return count
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
