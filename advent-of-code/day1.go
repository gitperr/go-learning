package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type elf struct {
	number               int
	caloriesCarriedTotal int
}

var elves = make([]elf, 0)

//var elves = make(map[int]elf)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberRegex := regexp.MustCompile(`\d+`)
	totalPerElf := 0
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line as needed
		if numberRegex.MatchString(line) {
			// Line contains a number, process it
			calorie, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			totalPerElf += calorie
			//fmt.Println(line)
		} else {
			newElf := elf{}
			newElf.number = i
			newElf.caloriesCarriedTotal = totalPerElf
			elves = append(elves, newElf)
			i++
			totalPerElf = 0
		}
		//fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to scan file: %v", err)
		return
	}
	//fmt.Println(elves)
	highestCalories := 0

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].caloriesCarriedTotal < elves[j].caloriesCarriedTotal
	})

	lastThree := elves[len(elves)-3:]
	fmt.Println(lastThree)
	totalOfLastThree := 0
	for _, elf := range lastThree {
		totalOfLastThree += elf.caloriesCarriedTotal
	}
	for _, elf := range elves {
		//fmt.Println(elf.number, elf.caloriesCarriedTotal)
		if elf.caloriesCarriedTotal >= highestCalories {
			highestCalories = elf.caloriesCarriedTotal
		}

	}
	fmt.Println(highestCalories)
	fmt.Println(totalOfLastThree)
}
