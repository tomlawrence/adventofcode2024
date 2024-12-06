package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day05/day05.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	rules := make(map[int][]int)
	updates := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		switch {
		case strings.Contains(scanner.Text(), "|"):
			rule := strings.Split(scanner.Text(), "|")
			page1, err1 := strconv.Atoi(rule[0])
			page2, err2 := strconv.Atoi(rule[1])
			if err1 != nil && err2 != nil {
				fmt.Printf("Error converting page numbers into rules: %v\n", err)
				continue
			}
			rules[page1] = append(rules[page1], page2)
		case strings.Contains(scanner.Text(), ","):
			pageNums := strings.Split(scanner.Text(), ",")
			update := []int{}
			for i := 0; i < len(pageNums); i++ {
				num, err := strconv.Atoi(pageNums[i])
				if err != nil {
					fmt.Printf("Error converting page number into int for update: %v\n", err)
					continue
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var part1, part2 int = sumMiddlePageNumbers(rules, updates)
	//fmt.Println("Valid updates middle sum: ", oldSumMiddlePageNumbers(rules, updates))
	fmt.Println("Part 1 - Sum of middle values of valid updates:", part1)
	fmt.Println("Part 2 - Sum of middle values of sorted invalid updates:", part2)
}

func sumMiddlePageNumbers(rules map[int][]int, updates [][]int) (int, int) {
	var sumPart1, sumPart2 int
	var cmp = func(a, b int) int { // Declare comparison function as defined by SortFunc
		for _, rule := range rules[a] { // For each rule for a
			if rule == b { // Check if page num is listed as a rule to come after a
				return -1 // Returning -1 meaning a comes before b (Go SortFunc docs)
			}
		}
		return 0 // No specific order (rule doesn't exist)
	}

	for _, update := range updates {
		if slices.IsSortedFunc(update, cmp) { // Check if update is sorted based on cmp func
			sumPart1 += update[len(update)/2]
		} else {
			slices.SortFunc(update, cmp)
			sumPart2 += update[len(update)/2]
		}
	}
	return sumPart1, sumPart2
}
