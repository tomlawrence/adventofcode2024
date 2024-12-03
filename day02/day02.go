package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day02/day02.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		lineNums := []int{}
		for i := 0; i < len(line); i++ {
			num, err := strconv.Atoi(line[i])
			if err == nil {
				lineNums = append(lineNums, num)
			}
		}
		reports = append(reports, lineNums)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	var safeReports int
	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			safeReports++
		}
	}
	fmt.Println("Safe reports:", safeReports)
}

func part2(reports [][]int) {
	var safeReports int
	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			safeReports++
		} else {
			for j := 0; j < len(reports[i]); j++ {
				newReport := make([]int, len(reports[i]))
				copy(newReport, reports[i])
				newReport = append(newReport[:j], newReport[j+1:]...)
				if isSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println("Safe reports:", safeReports)
}

func isSafe(report []int) bool {
	var safe bool = true
	var ascending bool

	for safe {
		for j := 1; j < len(report); j++ {
			diff := report[j] - report[j-1]
			if j == 1 {
				if diff > 0 {
					ascending = true
				}
			}
			if (diff == 0) || (ascending && diff <= 0) || (ascending && diff > 3) ||
				(!ascending && diff >= 0) || (!ascending && diff < -3) {
				safe = false
				break
			}
			if j == len(report)-1 {
				return true
			}
		}
	}
	return false
}
