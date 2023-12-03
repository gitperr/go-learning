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
	// parse a line and return bool (possible/impossible) and the game ID

	colorsAndLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	numberRegex := regexp.MustCompile("[0-9]+")
	// remove spaces
	spacesRemoved := strings.Replace(line, " ", "", -1)
	gameAndId := strings.Split(spacesRemoved, ":")             // Game 1:, Game 2:
	numberSlice := numberRegex.FindAllString(gameAndId[0], -1) // grab just game ID number e.g. 1, 2 etc.
	gameId, err := strconv.Atoi(numberSlice[0])
	if err != nil {
		fmt.Println("Error in conversion")
	}
	colorCubesString := gameAndId[1] // refers to the cubes in the game 1red,2green etc.
	colorCubesSplit := strings.Split(colorCubesString, ";")
	// take a round of color cubes
	for _, round := range colorCubesSplit {
		roundSplit := strings.Split(round, ",")
		// iterate over colors in the round
		// countAndColor: 1red, 1green etc...
		// so we need to still extract the count
		for _, countAndColor := range roundSplit {
			for color, limit := range colorsAndLimits {
				if strings.Contains(countAndColor, color) {
					// extract the count from countAndColor
					// 1, 2, 3 etc.
					numberSlice := numberRegex.FindAllString(countAndColor, -1)
					count, err := strconv.Atoi(numberSlice[0])
					if err != nil {
						fmt.Println("Error during conversion")
					}
					if count <= limit {
						continue
					} else {
						return -1
					}

				}

			}
		}
	}

	return gameId
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
		gameId := gameParser(line)
		if gameId != -1 {
			total += gameId
		}
	}
	fmt.Println(total)
}
