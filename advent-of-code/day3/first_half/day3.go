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

func returnDuplicateLetter(firstCompartment string, secondCompartment string) string {
	for _, letter := range firstCompartment {
		// if the letter is present in second compartment
		letter := string(letter)
		if strings.Contains(secondCompartment, letter) {
			return letter
		}
	}
	return "error"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		halfRucksackLength := len(rucksack) / 2
		// parse the rucksack into two compartments
		firstCompartment := rucksack[:halfRucksackLength]
		secondCompartment := rucksack[halfRucksackLength:]
		// find duplicate item, and assign it into
		// item struct, which will give it a prio as well
		// then put it into items list for later processing
		duplicateItem := returnDuplicateLetter(firstCompartment, secondCompartment)
		newItem := returnCreatedItem(duplicateItem, letters)
		itemsList = append(itemsList, newItem)
	}
	var sumOfPriorities int = 0
	for _, item := range itemsList {
		sumOfPriorities += item.priority
	}
	fmt.Println(sumOfPriorities)
}
