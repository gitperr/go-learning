package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func setFilePath() string {
	var filePath = flag.String("file", "", "File path")
	flag.Parse()
	log.Println("File path was set to", *filePath)
	if filePath != nil {

		return *filePath
	}
	return ""
}

func readCsvFile(filePath string) [][]string {
	if filePath == "" {
		log.Println("No file path given, defaulting to problems.csv")
		filePath = "problems.csv"
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as csv for "+filePath, err)
	}

	return records

}

func main() {
	message := "Hello, this is abi's quiz program"
	fmt.Println(message)
	timeLimit := flag.Int("limit", 30, "Set a time limit in seconds to complete the quiz")
	questionLimit := flag.Int("questionlimit", 30, "Set a question count.")
	filepath := setFilePath()
	records := readCsvFile(filepath)
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	correct_answers := 0
	//elapsed_time_seconds := 0
	answerChannel := make(chan string)
	i := 0
	for i < *questionLimit {
		question := records[i][0]
		correct_answer := records[i][1]
		i++
		fmt.Println("Question:", question)
		go func() {
			var answer strings
			fmt.Println("Your answer: ")
			fmt.Scanln(&answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			return
		case answer := <-answerChannel:
			if answer == correct_answer {
				correct_answers++
			}
		}
	}

	fmt.Printf("You have responded correctly %v/%v times!\n", correct_answers, *questionLimit)

}
