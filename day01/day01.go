package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day01/day01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		leftNum, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error converting left number: %v\n", err)
			continue
		}
		rightNum, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("Error converting right number: %v\n", err)
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	part1(leftList, rightList)
	part2(leftList, rightList)

}

func part1(leftList, rightList []int) {
	var totalDistance int

	for i := 0; i < len(leftList); i++ {
		totalDistance += absoluteDiff(leftList[i], rightList[i])
	}
	fmt.Println("Total Distance:", totalDistance)
}

func part2(leftList, rightList []int) {
	var similarityScore int
  counted := make(map[int]int)

  for i := 0; i < len(leftList); i++ {
    // Check for previously counted value
    if i > 0 && leftList[i] == leftList[i-1] {
      similarityScore += leftList[i] * counted[i]
      continue
    }

    var count int
    for j := 0; j < len(rightList); j++ {
      if leftList[i] == rightList[j] {
        count++
      }
      //Since array is sorted, stop looping rest of array once greater than current value
      if rightList[j] > leftList[i] {
        continue
      }
    }
    counted[leftList[i]] = count
    similarityScore += leftList[i] * count
  }
  fmt.Println("Similarity Score:", similarityScore)
}

func absoluteDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
