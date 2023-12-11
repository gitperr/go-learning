package main

import (
	"os"
	"strings"
	"testing"
)

func TestRealInput(t *testing.T) {
	path := "fixed_input.txt"
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	contentString := strings.ReplaceAll(string(content), "\r\n", "\n")
	contentString = strings.ReplaceAll(contentString, "\r", "\n")
	// Convert content to string and split it into lines
	lines := strings.Split(string(contentString), "\n")
	matrix := make([][]string, len(lines))
	//fmt.Println("line count:", len(lines))
	for i, line := range lines {
		//fmt.Println("line num:", i, "is:", line)
		matrix[i] = strings.Split(line, "")

	}
	result := matrixNumberDetector(matrix)
	expected := 526404

	if result != expected {
		t.Errorf("matrixNumberDetector(matrix) returned %d, expected %d", result, expected)
	}
}

func TestCustomInput(t *testing.T) {
	path := "test_input2.txt"
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	contentString := strings.ReplaceAll(string(content), "\r\n", "\n")
	contentString = strings.ReplaceAll(contentString, "\r", "\n")
	// Convert content to string and split it into lines
	lines := strings.Split(string(contentString), "\n")
	matrix := make([][]string, len(lines))
	//fmt.Println("line count:", len(lines))
	for i, line := range lines {
		//fmt.Println("line num:", i, "is:", line)
		matrix[i] = strings.Split(line, "")

	}
	result := matrixNumberDetector(matrix)
	expected := 4361

	if result != expected {
		t.Errorf("matrixNumberDetector(matrix) returned %d, expected %d", result, expected)
	}
}
