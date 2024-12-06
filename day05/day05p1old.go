package main

import (
	"fmt"
	"slices"
)

// Old function that I initally wrote for part 1 before I did a rewrite for part 2
func oldSumMiddlePageNumbers(rules map[int][]int, updates [][]int) int {
	var sum int
	for _, update := range updates {
		var valid bool = true
		for i := 0; i < len(update); i++ {
			pageNum := update[i]
			for _, rule := range rules[pageNum] {
				previousPages := update[:i]
				index := slices.Index(previousPages, rule)
				if index != -1 {
					valid = false
					break
				}
			}
		}
		if valid {
			fmt.Println("Valid list:", update)
		}
	}
	return sum
}
