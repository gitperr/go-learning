package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isSymbol(char string) bool {
	//fmt.Println("Checking if char", char, "contains special stuff. Result->", strings.Contains("*%&$#+-?@=/", char))
	return strings.Contains("*%&$#+-?@=/", char)
}

func lookLeft(matrix [][]string, i int, x int) string {
	//fmt.Println("values passed:", i, x)
	if matrix[i][x] == "X" {
		return "end"
	}
	return matrix[i][x-1]
}

func lookTopLeft(matrix [][]string, i int, x int) string {
	//fmt.Println("values passed:", i, x)
	if i == 0 || x == 0 {
		return "end"
	}
	return matrix[i-1][x-1]
}

func lookTop(matrix [][]string, i int, x int) string {
	if i == 0 || x == len(matrix)-1 {
		return "end"
	}
	return matrix[i-1][x]
}

func lookTopRight(matrix [][]string, i int, x int) string {
	if i == 0 || x == len(matrix)-1 {
		return "end"
	}
	return matrix[i-1][x+1]
}

func lookBotLeft(matrix [][]string, i int, x int) string {
	if i == 0 || x == len(matrix)-1 || x == 0 {
		return "end"
	}
	return matrix[i+1][x-1]
}

func lookBot(matrix [][]string, i int, x int) string {
	if i == 0 || x == len(matrix)-1 {
		return "end"
	}
	return matrix[i+1][x]
}

func lookBotRight(matrix [][]string, i int, x int) string {
	if i == 0 || x == len(matrix)-1 || matrix[i][x] == "X" {
		return "end"
	}
	return matrix[i+1][x+1]
}

func convertToInt(columnString string) int {
	//fmt.Println("Number passed is:", columnString)
	if columnString != "" {
		number, err := strconv.Atoi(columnString)
		if err != nil {
			panic(err)
		}
		return number
	} else {
		return 0
	}
}

func doesPerimeterContainSymbol(columnPerimeter []string) bool {
	// check the perimeter for symbol
	// if any symbol is found, instantly return true
	// otherwise, always assume we could not find a symbol and return false
	for _, value := range columnPerimeter {
		if isSymbol(value) {
			//fmt.Println("Checking value", value)
			//fmt.Println("Perimeter contains symbol! Returning true")
			return true
		}
	}
	return false
}

func matrixNumberDetector(matrix [][]string) int {
	//fmt.Println("matrix given:", matrix)
	total := 0
	symbolFound := false
	numberRegex := regexp.MustCompile("[0-9]+")
	//symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	foundNumber := ""
	for i, row := range matrix {
		for x, column := range row {
			// set column perimeter to empty for clean lookup
			var columnPerimeter []string
			// if we found a number, start concatenating it
			if matrix[i][x] != "X" {
				// fmt.Println("top left:", lookTopLeft(matrix, i, x))
				// fmt.Println("top:", lookTop(matrix, i, x))
				// fmt.Println("top right:", lookTopRight(matrix, i, x))
				// fmt.Println("bot left:", lookBotLeft(matrix, i, x))
				// fmt.Println("bot:", lookBot(matrix, i, x))
				// fmt.Println("bot right:", lookBotRight(matrix, i, x))
				//fmt.Println("Current foundNumber is", foundNumber)
				//fmt.Println("columnValue is", column)
				if numberRegex.FindAllString(column, -1) != nil {
					foundNumber = foundNumber + column
					// start the look around and construct a perimeter slice
					columnPerimeter = append(columnPerimeter, lookLeft(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookTopLeft(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookTop(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookTopRight(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookBotLeft(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookBot(matrix, i, x))
					columnPerimeter = append(columnPerimeter, lookBotRight(matrix, i, x))
					// check for symbol in the perimeter
					//fmt.Println("columnPerimeter is:", columnPerimeter, "starting lookup!")
					if symbolFound {
						if len(foundNumber) > 1 {
							continue
						} else {
							symbolFound = doesPerimeterContainSymbol(columnPerimeter)
						}
					} else {
						symbolFound = doesPerimeterContainSymbol(columnPerimeter)
					}
				} else {
					//fmt.Println("Current foundNumber after if is", foundNumber, "columnValue is:", column)
					if column == "." || column == "X" {
						if !symbolFound {
							//fmt.Println("Symbol was not found, setting", foundNumber, "to", "empty string")
							foundNumber = ""
						} else {
							number := convertToInt(foundNumber)
							//fmt.Println("Adding", number, "to", total)
							total = number + total
							foundNumber = ""
						}
						symbolFound = false
					} else {
						if !symbolFound {
							symbolFound = true
						}
						number := convertToInt(foundNumber)
						//fmt.Println("Adding", number, "to", total)
						total = number + total
						foundNumber = ""
						if number != 0 {
							symbolFound = false
						}
					}
				}
			} else {
				// if we have some number left at the end, we need to add it if symbol was found
				if i == len(matrix)-1 && symbolFound {
					number := convertToInt(foundNumber)
					fmt.Println("Adding", number, "to", total)
					total = number + total
					foundNumber = ""
				}

			}
		}
	}
	return total
}
