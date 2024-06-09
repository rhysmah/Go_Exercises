package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type UserResponse struct {
	TotalQuestions   int
	CorrectAnswers   int
	IncorrectAnswers int
	Percentage       float64
}

// processQuestionsAndAnswers reads a CSV file containing questions and answers
// and returns a map where keys are questions, values are answers.
//
// The CSV file is expected to have each question-answer pair in a separate row,
// with the question in the first column and the answer in the second column.
//
// Parameters:
// - filename: A string representing the path to the CSV file to be read.
//
// Returns:
// - A map[string]string where each key is a question, each value the answer
// - An error if there was any issue opening the file or parsing its contents.
func processQuestionsAndAnswers(filename string) (map[string]string, error) {
	questionsAndAnswers := make(map[string]string)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	parseError := parseCSVFile(file, questionsAndAnswers)
	if parseError != nil {
		return nil, parseError
	}

	return questionsAndAnswers, nil
}

// parseCSVFile reads a CSV file and populates a provided map with questions and answers.
//
// This function reads each record from the given CSV file and stores it in the provided map.
// The first column of each record is treated as the question, and the second column is treated
// as the answer. Both the question and answer are trimmed of leading and trailing whitespace
// before being stored in the map.
//
// Parameters:
// - file: A pointer to an os.File representing the open CSV file to be read.
// - quizQA: A map[string]string where the questions and answers will be stored.
//
// Returns:
//   - An error if there was any issue reading the CSV file. If the end of the file is reached,
//     or if individual records have errors (which are logged), the function will not return an error.
//
// Logging:
// - The function logs a message when the end of the file is reached.
// - It logs errors encountered while reading records but continues processing the remaining records.
// - It logs invalid records (those with fewer than two fields) and continues processing.
func parseCSVFile(file *os.File, quizQA map[string]string) error {
	r := csv.NewReader(file)

	for {
		record, err := r.Read()

		if err == io.EOF {
			log.Print("Reached end of file")
			break
		}
		if err != nil {
			log.Printf("Error reading record: %v", err)
			continue
		}
		if len(record) < 2 {
			log.Printf("Invalid record (less than two fields): %v", record)
			continue
		}
		quizQA[strings.TrimSpace(record[0])] = strings.TrimSpace(record[1])
	}
	return nil
}

func main() {

	questionsAndAnswers, err := processQuestionsAndAnswers("questions.csv")
	if err != nil {
		log.Fatalf("Error processing questions and answers: %v", err)
	}
	fmt.Println(questionsAndAnswers)
}
