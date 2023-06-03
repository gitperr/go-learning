package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type round struct {
	elfShape    string
	playerShape string
}

var rounds = make([]round, 0)

func roundResult(round round) string {
	elfShape := round.elfShape
	playerShape := round.playerShape

	if elfShape == playerShape {
		return "draw"
	}
	if elfShape == "A" && playerShape != "C" {
		return "playerWin"
	}
	if elfShape == "B" && playerShape != "A" {
		return "playerWin"
	}
	if elfShape == "C" && playerShape != "B" {
		return "playerWin"
	}
	return "elfWin"
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//numberRegex := regexp.MustCompile(`\d+`)
	//totalPerElf := 0
	//i := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line as needed
		// Split the line into two letters
		letters := strings.Split(line, "")

		// Print the two letters separated by a space
		//fmt.Println(letters[0], letters[1])
		elfShape := letters[0]
		playerShape := letters[1]
		if playerShape == "X" {
			playerShape = "A"
		}
		if playerShape == "Y" {
			playerShape = "B"
		}
		if playerShape == "Z" {
			playerShape = "C"
		}
		newRound := round{elfShape, playerShape}
		rounds = append(rounds, newRound)

	}
	//fmt.Println(rounds)
	playerTotalScore := 0
	shapes := []string{"A", "B", "C"}
	for _, round := range rounds {
		//fmt.Println(round.elfShape)
		for index, shape := range shapes {
			if shape == round.playerShape {
				playerTotalScore += +index + 1
			}
		}
		roundResult := roundResult(round)
		if roundResult == "draw" {
			playerTotalScore += 3
		}
		if roundResult == "playerWin" {
			playerTotalScore += 6
		}
	}
	fmt.Println(playerTotalScore)
}
