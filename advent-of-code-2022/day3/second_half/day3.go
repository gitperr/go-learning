package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type item struct {
	letter   string
	priority int
}

var itemsList = make([]item, 0)

func returnDuplicateLetter(rucksacks []string) string {
	commonLetter := ""
	firstRucksack := rucksacks[0]

	for _, char := range firstRucksack {
		letter := string(char)
		isCommon := true

		// check other rucksacks for the same letter
		for _, otherRucksack := range rucksacks[1:] {
			if !strings.Contains(otherRucksack, letter) {
				isCommon = false
				break
			}
		}

		if isCommon {
			commonLetter = letter
			break
		}
	}
	return commonLetter
}

func returnCreatedItem(letter string, prioList []string) item {
	var newItem = item{}
	for priority, prioLetter := range prioList {
		if letter == prioLetter {
			// create a new item with the letter and priority
			// priority gets a +1 because indexing starts from 0
			// but prioritization starts from 1
			newItem = item{
				letter:   letter,
				priority: priority + 1,
			}
		}
	}
	return newItem
}

func main() {
	// Create an empty slice
	letters := []string{}
	// Append lowercase letters
	for c := 'a'; c <= 'z'; c++ {
		letters = append(letters, string(c))
	}
	// Append uppercase letters
	for c := 'A'; c <= 'Z'; c++ {
		letters = append(letters, string(c))
	}
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	// elves are in groups of three
	groupSize := 3
	elfGroupRucksacks := make([]string, 0, groupSize)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		elfGroupRucksacks = append(elfGroupRucksacks, rucksack)
		// when the group is full AKA 3 elves/rucksacks are there
		// find and return the duplicate item and
		// reset the group to empty, so we can put new elves
		if len(elfGroupRucksacks) == groupSize {
			duplicateItem := returnDuplicateLetter(elfGroupRucksacks)
			newItem := returnCreatedItem(duplicateItem, letters)
			itemsList = append(itemsList, newItem)
			elfGroupRucksacks = make([]string, 0, groupSize)
		}
	}
	var sumOfPriorities int = 0
	for _, item := range itemsList {
		sumOfPriorities += item.priority
	}
	fmt.Println(sumOfPriorities)
}
