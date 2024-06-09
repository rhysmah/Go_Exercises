package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type User struct {
	UserName         string
	TotalQuestions   int
	CorrectAnswers   int
	IncorrectAnswers int
	Percentage       float64
}

// Pass in a pointer because we're modifying the struct
func (u *User) calculatePercentage() {
	if u.TotalQuestions <= 0 {
		u.Percentage = 0.0
	}
	u.Percentage = float64(u.CorrectAnswers/u.TotalQuestions) * 100
}

func (u *User) printResults() {

	fmt.Println("*****************************")
	fmt.Println("Username: ", u.UserName)
	fmt.Println("*****************************")
	fmt.Println("# Questions: ", u.TotalQuestions)
	fmt.Println("Correct: ", u.CorrectAnswers)
	fmt.Println("Incorrect: ", u.IncorrectAnswers)
	fmt.Println("Percentage: ", u.Percentage)
}

// processQuestionsAndAnswers reads a CSV file containing questions and answers
// and returns a map where keys are questions, values are answers.
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

func runQuiz(quizQA map[string]string) {
	var userResponse User

	userResponse.TotalQuestions = len(quizQA)

	fmt.Print("Please enter your name: ")
	fmt.Scanln(&userResponse.UserName)

	for question, answer := range quizQA {
		fmt.Println(question)

		var userAnswer string
		fmt.Scanln(&userAnswer)

		if userAnswer == answer {
			fmt.Println("Correct!")
			userResponse.CorrectAnswers++
		} else {
			fmt.Println("Incorrect!")
			userResponse.IncorrectAnswers++
		}
	}
	userResponse.calculatePercentage()
	userResponse.printResults()

}

func main() {

	quizQA, err := processQuestionsAndAnswers("questions.csv")
	if err != nil {
		log.Fatal("Error processing questions and answers: ", err)
	}
	if len(quizQA) <= 0 {
		log.Fatal("No questions found in the quiz. Exiting...")
	}
	runQuiz(quizQA)
}
