package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type calibrationValue struct {
	firstDigit  string
	secondDigit string
	total       string
}

func main() {
	filePath := "../input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	numberRegex := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	grandTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		if numberRegex.MatchString(line) {
			// get numbers from matching lines
			// they look like this now:
			// 3627837 7 (first number *space* second number etc.)
			// 5 1 or 1 1 ...
			numberSlice := numberRegex.FindAllString(line, -1)
			// need to handle some cases where the any of the numbers
			// has many digits, we are just interested in the first
			// digit of first one, and last digit of last one
			calibration := getCalibrationValue(numberSlice)
			totalCalibrationInt, err := strconv.Atoi(calibration.total)
			if err != nil {
				fmt.Println("Error during conversion to int")
			}
			grandTotal = grandTotal + totalCalibrationInt
		}
	}
	fmt.Println(grandTotal)
}

func getCalibrationValue(numberSlice []string) calibrationValue {
	// get the numberSlice and return first and last digit of it
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
	return newCalibrationValue
}
