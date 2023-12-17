package main

import (
	"math"
	"os"
	"strings"
)

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func countWinningCards(numToLook string, cardList []string) int {
	occurrence := 0
	limit := len(cardList)
	for i := 0; i < limit; i++ {
		if cardList[i] == numToLook && numToLook != "" {
			//fmt.Println(cardList[i], "is equal to", numToLook, "adding!")
			occurrence += 1
		}
	}
	return occurrence
}

func resultIs(filePath string) int {
	//filePath := "../test_input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	total := 0
	contentString := strings.ReplaceAll(string(content), "\r\n", "\n")
	contentString = strings.ReplaceAll(contentString, "\r", "\n")
	// Convert content to string and split it into lines
	lines := strings.Split(string(contentString), "\n")
	for _, line := range lines {
		line := strings.Split(string(line), " ")
		//fmt.Println(line[2:(len(line))])
		var separatorIndex int
		if len(line) > 0 {
			separatorIndex = sliceIndex(len(line), func(i int) bool { return line[i] == "|" })
		}
		//fmt.Printf("SeparatorIndex is %s in line %s", string(separatorIndex), line)
		winningNumbers := line[2:separatorIndex]
		ourCards := line[separatorIndex+1:]
		wonCount := 0
		for _, winningNumber := range winningNumbers {
			wonCount += countWinningCards(winningNumber, ourCards)
		}
		//fmt.Println("Winning numbers:", winningNumbers, "our cards:", ourCards, "we have", wonCount, "matches")
		powered := math.Pow(2.0, float64(wonCount-1))
		total += int(powered)
	}
	return total
}
