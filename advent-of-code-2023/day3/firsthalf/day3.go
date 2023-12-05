package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func matrixNumberDetector(matrix [][]string) int {
	//fmt.Println("matrix given:", matrix)
	total := 0
	symbolFound := false
	numberRegex := regexp.MustCompile("[0-9]+")
	//symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	symbolRegex := strings.Contains("$*#+@")
	foundNumber := ""

	for i, row := range matrix {
		for x, column := range row {
			// if we found a number, start concatenating it
			if numberRegex.FindAllString(column, -1) != nil {
				foundNumber = foundNumber + column
				// matrix[i][x] is our current location which is a number
				// if we haven't already found a symbol, start the lookup
				fmt.Println("row:", i, "column:", x, "Matrix length:", len(matrix))
				fmt.Println("foundnumber:", foundNumber, "total:", total, "currently at character", matrix[i][x])
				if !symbolFound {
					// if we are NOT at the first or last row
					if i > 0 && x > 0 && i < len(matrix)-2 {
						//if i-1 > -1 && x-1 > -1 && i < len(matrix)-2 {
						//fmt.Println("Trying to access", matrix[i-1][x-1])
						// look diagonally (top left)
						//fmt.Println("Trying matrix[i-1][x-1]. Current i and x", i, x)
						if symbolRegex.FindAllString(matrix[i][x-1], -1) != nil && !symbolFound {
							// if we found symbol, set flag
							//fmt.Println("Accessed matrix[i-1][x-1]")
							symbolFound = true
						}
						if symbolRegex.FindAllString(matrix[i][x+1], -1) != nil && !symbolFound {
							// if we found symbol, set flag
							//fmt.Println("Accessed matrix[i-1][x-1]")
							symbolFound = true
						}
						if symbolRegex.FindAllString(matrix[i-1][x-1], -1) != nil && !symbolFound {
							// if we found symbol, set flag
							//fmt.Println("Accessed matrix[i-1][x-1]")
							symbolFound = true
						}
						// look above
						//fmt.Println("Trying matrix[i-1][x]. Current i and x", i, x)
						if symbolRegex.FindAllString(matrix[i-1][x], -1) != nil && !symbolFound {
							//fmt.Println("Accessed matrix[i-1][x]")
							symbolFound = true
						}
						// look diagonally (top right)
						//fmt.Println("Trying matrix[i-1][x+1]. Current i and x", i, x)
						if symbolRegex.FindAllString(matrix[i-1][x+1], -1) != nil && !symbolFound {
							//fmt.Println("Accessed matrix[i-1][x+1]")
							symbolFound = true
						}
						// look diagonally (bottom left)
						//fmt.Println("Trying matrix[i+1][x-1]. Current i and x", i, x, "Current character is", matrix[i][x])
						if symbolRegex.FindAllString(matrix[i+1][x-1], -1) != nil && !symbolFound {
							//fmt.Println("Accessed matrix[i+1][x-1]")
							symbolFound = true
						}
						// look below
						//fmt.Println("Trying matrix[i+1][x]. Current i and x", i, x)
						if symbolRegex.FindAllString(matrix[i+1][x], -1) != nil && !symbolFound {
							//fmt.Println("Accessed matrix[i+1][x]")
							symbolFound = true
						}
						// look diagonally (bottom right)
						//fmt.Println("Trying matrix[i+1][x+1]. Current i and x", i, x)
						if symbolRegex.FindAllString(matrix[i+1][x+1], -1) != nil && !symbolFound {
							//fmt.Println("Accessed matrix[i+1][x+1]")
							symbolFound = true
						}
						// if we are at first or last row, we can only look down OR up (NOT BOTH!) depending on where we are
					} else {
						// if we are at the first row and first column
						if i == 0 && x == 0 {
							if symbolRegex.FindAllString(matrix[i+1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i+1][x+1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
						// if we are at first row, and some other column
						if i == 0 && x != 0 {
							if symbolRegex.FindAllString(matrix[i][x+1], -1) != nil && !symbolFound {
								// if we found symbol, set flag
								//fmt.Println("Accessed matrix[i-1][x-1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i][x-1], -1) != nil && !symbolFound {
								// if we found symbol, set flag
								//fmt.Println("Accessed matrix[i-1][x-1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i+1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i+1][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x-1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i+1][x+1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
						// if we are at first row, and last column, we cannot look right, or up
						if i == 0 && x == -1 {
							if symbolRegex.FindAllString(matrix[i+1][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x-1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i+1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
						// if we are at last row, and first column, we cannot look left, or down
						if i == len(matrix)-2 && x == 0 {
							if symbolRegex.FindAllString(matrix[i-1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i-1][x+1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
						// if we are at last row and last column, we cannot look right, or down
						if i == len(matrix)-2 && x == -1 {
							if symbolRegex.FindAllString(matrix[i-1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i-1][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
						// if we are on some other column on last row, look left and right, above...
						if i == len(matrix)-2 {
							if symbolRegex.FindAllString(matrix[i-1][x], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i-1][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i-1][x+1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i][x+1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
							if symbolRegex.FindAllString(matrix[i][x-1], -1) != nil && !symbolFound {
								//fmt.Println("Accessed matrix[i+1][x+1]")
								symbolFound = true
							}
						}
					}
				}
			} else {
				// if it was not a dot, we are in luck, because we found a symbol next to the number.
				// can break the search here and return the number
				if column != "." && len(foundNumber) > 0 {
					// reset symbol flag
					symbolFound = false
					// add the number to total
					number, err := strconv.Atoi(foundNumber)
					if err != nil {
						fmt.Println("Failed to convert", foundNumber, "to int")
					}
					total += number
					// reset back to empty string
					foundNumber = ""
				} else {
					// if we found a dot, check if symbol flag was set, add to total
					// etc.
					//fmt.Println("Found dot, resetting.")
					if symbolFound && len(foundNumber) > 0 {
						symbolFound = false
						number, err := strconv.Atoi(foundNumber)
						if err != nil {
							fmt.Println("Error in conversion")
						}
						total += number
						foundNumber = ""
					} else {
						foundNumber = ""
					}

				}
			}
		}
	}
	return total
}

func main() {
	//total := 0
	filePath := "../input.txt"
	// file, err := os.Open(filePath)
	// Read content from the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	//line := scanner.Text()
	// Convert content to string and split it into lines
	lines := strings.Split(string(content), "\n")
	// Create a 2D matrix
	matrix := make([][]string, len(lines)+1)
	for i, line := range lines {
		// Split each line into individual characters
		matrix[i] = strings.Split(line, "")
	}
	total := matrixNumberDetector(matrix)
	fmt.Println(total)
}
