package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func setTimerSeconds() time.Duration {
	var timer = flag.Int("timer", 30, "Set countdown timer")
	flag.Parse()
	log.Printf("Countdown timer was set to %v", *timer)
	return time.Duration(*timer) * time.Second
}

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
	time_limit_seconds := setTimerSeconds()
	fmt.Println(time_limit_seconds)
	//start := time.Now()
	filepath := setFilePath()
	records := readCsvFile(filepath)
	total_questions := 10
	wrong_answers := 0
	//elapsed_time_seconds := 0
	i := 0
	// tick once per second
	for range time.Tick(1 * time.Second) {
		time_limit_seconds -= 1 * time.Second
		fmt.Println(time_limit_seconds)
		//end := time.Now()
		//remainingTime := end.Sub(time_limit_seconds)
		//fmt.Println(remainingTime)
		question := records[i][0]
		correct_answer := records[i][1]
		i++
		fmt.Println("Question:", question)
		fmt.Println("Your answer: ")
		var answer string
		fmt.Scanln(&answer)
		if answer != correct_answer {
			wrong_answers++
		}
	}

	fmt.Printf("You have responded correctly %v/%v times!\n", total_questions-wrong_answers, total_questions)
	//fmt.Printf("Total time spent was %v seconds!", //elapsed_time_seconds)

}
