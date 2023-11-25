package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack []interface{}

func splitIntoColumns(line string, columnSize int) []string {
	var columns []string

	// Pad the line with spaces to ensure it can be divided evenly into columns
	line = line + strings.Repeat(" ", columnSize-len(line)%columnSize)

	// Split the line into columns
	for i := 0; i < len(line); i += columnSize {
		column := line[i : i+columnSize]
		columns = append(columns, column)
	}
	return columns
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func main() {
	//stack := Stack{}

	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into columns of three characters
		columns := splitIntoColumns(line, 3)

		// Print the columns
		fmt.Println(columns)
		//fmt.Println(parsedLine)
		i++
		if i == 10 {
			break
		}
		// fmt.Println(splitLine)

		// stack.Push(10)
		// stack.Push(20)
		// stack.Push(30)

		// for !stack.IsEmpty() {
		// 	value, _ := stack.Pop()
		// 	fmt.Println(value)
		// }
	}
}
