package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "../input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	contentString := strings.ReplaceAll(string(content), "\r\n", "\n")
	contentString = strings.ReplaceAll(contentString, "\r", "\n")
	// Convert content to string and split it into lines
	lines := strings.Split(string(contentString), "\n")

	// Process each line and wrap it with "X"
	for i, line := range lines {
		lines[i] = "X" + line + "X"
	}
	// Add lines full of "X" before and after the first and last lines
	prefix := strings.Repeat("X", len(lines[1]))
	lines = append([]string{prefix}, lines...)
	lines = append(lines, prefix)
	outputContent := strings.Join(lines, "\n")
	//fmt.Println(outputContent)
	// Write the modified content to a new file
	fixedPath := "fixed_input.txt"
	err = os.WriteFile(fixedPath, []byte(outputContent), 0644)
	if err != nil {
		panic(err)
	}

	content, err = os.ReadFile(fixedPath)
	if err != nil {
		panic(err)
	}
	contentString = strings.ReplaceAll(string(content), "\r\n", "\n")
	contentString = strings.ReplaceAll(contentString, "\r", "\n")
	// Convert content to string and split it into lines
	lines = strings.Split(string(contentString), "\n")
	matrix := make([][]string, len(lines))
	//fmt.Println("line count:", len(lines))
	for i, line := range lines {
		//fmt.Println("line num:", i, "is:", line)
		matrix[i] = strings.Split(line, "")

	}
	result := matrixNumberDetector(matrix)
	fmt.Println(result)
}
