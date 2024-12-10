package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type File struct {
	Data int
	Size int
}

func main() {
	file, err := os.ReadFile("day09/day09.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	diskString := strings.Split(string(file), "")
	var diskMap []int
	for _, char := range diskString {
		number, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
			continue
		}
		diskMap = append(diskMap, number)
	}
	fmt.Println("Part 1 - Filesystem Checksum:", part1(diskMap))
	fmt.Println("Part 2 - Filesystem Checksum:", part2(diskMap))
}

func part1(diskMap []int) int {
	expandedDisk := expandDisk(diskMap)
	defraggedDisk := defragDisk(expandedDisk, false)
	return calculateChecksum(defraggedDisk)
}

func part2(diskMap []int) int {
	expandedDisk := expandDisk(diskMap)
	defraggedDisk := defragDisk(expandedDisk, true)
	return calculateChecksum(defraggedDisk)
}

func expandDisk(disk []int) []File {
	var expandedDisk []File
	for i, number := range disk {
		if number > 0 {
			if i%2 == 0 {
				expandedDisk = append(expandedDisk, File{Data: i / 2, Size: number})
			} else {
				expandedDisk = append(expandedDisk, File{Data: -1, Size: number})
			}
		}
	}
	return expandedDisk
}

func findLeftmostFreeSpace(disk []File, requiredSize int) int {
	for i := 0; i < len(disk); i++ {
		if disk[i].Data == -1 && (requiredSize == 0 || disk[i].Size >= requiredSize) {
			return i
		}
	}
	return -1
}

func findLastFile(disk []File, lastFileIndex int) int {
	for i := lastFileIndex; i >= 0; i-- {
		if disk[i].Data != -1 {
			return i
		}
	}
	return -1
}

func defragDisk(disk []File, wholeFilesOnly bool) []File {
	lastFileIndex := len(disk) - 1
	for lastFileIndex >= 0 {
		j := findLastFile(disk, lastFileIndex)
		requiredSize := 0
		if wholeFilesOnly {
			requiredSize = disk[j].Size
		}
		i := findLeftmostFreeSpace(disk, requiredSize)
		if i == -1 || i >= j {
			lastFileIndex = j - 1
			continue
		}
		disk = moveFile(disk, i, j)
		//printDisk(disk)
		lastFileIndex = len(disk) - 1
	}
	return disk
}

func moveFile(disk []File, freeSpaceIndex int, lastFileIndex int) []File {
	freeSpace := disk[freeSpaceIndex]
	lastFile := disk[lastFileIndex]
	diff := freeSpace.Size - lastFile.Size
	switch {
	case diff > 0: // File fits with space remaining
		disk[freeSpaceIndex] = lastFile
		disk[lastFileIndex] = File{Data: -1, Size: lastFile.Size}
		disk = slices.Insert(disk, freeSpaceIndex+1, File{Data: -1, Size: diff})
	case diff == 0: // File fits exactly
		disk[freeSpaceIndex] = lastFile
		disk[lastFileIndex] = freeSpace
	case diff < 0: // File is larger than space
		disk[freeSpaceIndex].Data = disk[lastFileIndex].Data
		disk[lastFileIndex].Size -= disk[freeSpaceIndex].Size
		disk = slices.Insert(disk, lastFileIndex+1, File{Data: -1, Size: freeSpace.Size})
	}
	return disk
}

func calculateChecksum(disk []File) int {
	var checksum int
	var position int
	for _, file := range disk {
		if file.Data != -1 {
			for i := 0; i < file.Size; i++ {
				checksum += position * file.Data
				position++
			}
		} else {
			position += file.Size
		}
	}
	return checksum
}

func printDisk(disk []File) {
	for _, file := range disk {
		if file.Data == -1 {
			fmt.Printf("%v", strings.Repeat(".", file.Size))
		} else {
			fmt.Printf("%v", strings.Repeat(strconv.Itoa(file.Data), file.Size))
		}
	}
	fmt.Printf("\n")
}
