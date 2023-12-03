package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func gameParser(line string) int {
	// parse a line and return power of minimum required color cube counts
	// this slice will be updated with the highest available count in the round of color cubes
	colorsAndHighestCount := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	numberRegex := regexp.MustCompile("[0-9]+")
	// remove spaces
	spacesRemoved := strings.Replace(line, " ", "", -1)
	gameAndId := strings.Split(spacesRemoved, ":") // Game 1:, Game 2:
	colorCubesString := gameAndId[1]               // refers to the cubes in the game 1red,2green etc.
	colorCubesSplit := strings.Split(colorCubesString, ";")
	// take a round of color cubes
	for _, round := range colorCubesSplit {
		roundSplit := strings.Split(round, ",")
		// iterate over colors in the round
		// countAndColor: 1red, 1green etc...
		// so we need to still extract the count
		for _, countAndColor := range roundSplit {
			for color, highestCount := range colorsAndHighestCount {
				if strings.Contains(countAndColor, color) {
					// extract the count from countAndColor
					// 1, 2, 3 etc.
					numberSlice := numberRegex.FindAllString(countAndColor, -1)
					count, err := strconv.Atoi(numberSlice[0])
					if err != nil {
						fmt.Println("Error during conversion")
					}
					if count > highestCount {
						colorsAndHighestCount[color] = count
					}

				}

			}
		}
	}
	power := colorsAndHighestCount["red"] * colorsAndHighestCount["blue"] * colorsAndHighestCount["green"]
	return power
}

func main() {
	total := 0
	filePath := "../input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		power := gameParser(line)
		total += power
	}
	fmt.Println(total)
}
