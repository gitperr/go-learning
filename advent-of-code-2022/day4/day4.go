package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rangesOverlap(start1, end1, start2, end2 int) bool {
	return end1 >= start2 && end2 >= start1
}

func splitAndReturnRangeNums(parts string) (int, int) {
	// example data structure:
	// [2-4]
	// then split each part from dash (-)
	// so each part is like this
	// [2 4]
	dashSplit := strings.Split(parts, "-")
	num1, err := strconv.Atoi(dashSplit[0])
	num2, err := strconv.Atoi(dashSplit[1])
	if err != nil {
		fmt.Printf("Error in converting to int")
	}
	return num1, num2

}

func createAndReturnRange(start int, end int) []int {
	var rangeToReturn []int
	for i := start; i <= end; i++ {
		rangeToReturn = append(rangeToReturn, i)
	}
	return rangeToReturn
}

func returnLongerAndShorterRanges(range1, range2 []int) ([]int, []int) {
	var longerRange []int
	var shorterRange []int
	if len(range1) > len(range2) {
		longerRange = range1
		shorterRange = range2
	} else {
		longerRange = range2
		shorterRange = range1
	}

	return longerRange, shorterRange

}
func isSubset(range1, range2 []int) bool {
	set := make(map[int]bool)
	for _, num := range range2 {
		set[num] = true
	}

	for _, num := range range1 {
		if !set[num] {
			return false
		}
	}

	return true
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	//fullyContained := 0
	overlap := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		assignments := scanner.Text()
		parts := strings.Split(assignments, ",")
		firstPart := parts[0]
		secondPart := parts[1]
		rangeStart, rangeEnd := splitAndReturnRangeNums(firstPart)
		//firstRange := createAndReturnRange(rangeStart, rangeEnd)
		start, end := splitAndReturnRangeNums(secondPart)
		//secondRange := createAndReturnRange(start, end)
		//longerRange, shorterRange := returnLongerAndShorterRanges(firstRange, secondRange)
		if rangesOverlap(rangeStart, rangeEnd, start, end) == true {
			overlap++
		}
		// if isSubset(shorterRange, longerRange) == true {
		// 	fullyContained++
		// }
	}
	fmt.Println(overlap)
}
