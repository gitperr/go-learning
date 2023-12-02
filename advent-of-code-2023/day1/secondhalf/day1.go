package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type calibrationValue struct {
	firstDigit  string
	secondDigit string
	total       string
}

func getCalibrationValue(numberSlice []string) calibrationValue {
	// get the numberSlice and return first and last digit of it
	//fmt.Println("slice given:", numberSlice)
	newCalibrationValue := calibrationValue{}
	iterators := []int{0, len(numberSlice) - 1}
	for _, i := range iterators {
		if i == 0 {
			firstDigitOfFirstNumber := string(numberSlice[i][0])
			newCalibrationValue.firstDigit = firstDigitOfFirstNumber
		}
		if i == len(numberSlice)-1 {
			lastDigitOfLastNumber := string(numberSlice[i][len(numberSlice[i])-1])
			newCalibrationValue.secondDigit = lastDigitOfLastNumber
		}
	}
	newCalibrationValue.total = newCalibrationValue.firstDigit + newCalibrationValue.secondDigit
	//fmt.Println("returning calibration value:", newCalibrationValue)
	return newCalibrationValue
}

func main() {
	// Number spellings map
	numberSpellings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	filePath := "../input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grandTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Number slice
		var numberSlice []string
		for index, char := range line {
			if unicode.IsDigit(char) {
				numberSlice = append(numberSlice, string(char))
			} else {
				for spelling, number := range numberSpellings {
					if strings.HasPrefix(line[index:], spelling) {
						//fmt.Println("found", spelling, "in", line, "appending", number, "to slice")
						numberSlice = append(numberSlice, string(number))
					}
					//numberSlice = append(numberSlice, " ")
				}
			}
		}
		calibration := getCalibrationValue(numberSlice)
		//fmt.Println("line:", line, "first digit:", calibration.firstDigit, "second digit", calibration.secondDigit, "concat:", calibration.total)
		totalCalibrationInt, err := strconv.Atoi(calibration.total)
		if err != nil {
			fmt.Println("Error during conversion to int")
		}
		grandTotal = grandTotal + totalCalibrationInt
	}
	fmt.Println(grandTotal)

}
