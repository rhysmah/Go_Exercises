package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	UserAnswer     string
	TotalQuestions int
	QuestionNumber int
	CorrectAnswers int
	Questions      map[string]string
}

// checkAnswer compares the user's answer with the correct answer.
func (q *Quiz) checkAnswer(question string, answer string) (bool, string) {
	if strings.TrimSpace(strings.ToLower(answer)) == question {
		return true, "Correct!"
	}
	return false, "Incorrect!"
}

// run executes the quiz, asking each question and checking answers.
func (q *Quiz) run(quizTime *time.Timer) {
	if len(q.Questions) == 0 {
		fmt.Println("No questions found in quiz data.")
		return
	}

	q.TotalQuestions = len(q.Questions)
	q.QuestionNumber = 1

	for question, answer := range q.Questions {

		select {
		case <-quizTime.C:
			fmt.Println("Time's up!")
			fmt.Printf("You answered %d out of %d questions correctly.\n",
				q.CorrectAnswers,
				q.TotalQuestions,
			)
			return

		default:
			fmt.Printf("Question %d: %s\n", q.QuestionNumber, question)
			q.QuestionNumber++

			fmt.Print("Your answer: ")
			fmt.Scanln(&q.UserAnswer)

			isCorrect, response := q.checkAnswer(answer, q.UserAnswer)
			fmt.Printf("%s\n\n", response)
			if isCorrect {
				q.CorrectAnswers++
			}
		}
	}
	fmt.Printf("You answered %d out of %d questions correctly.\n",
		q.CorrectAnswers,
		q.TotalQuestions,
	)
}

// processQuizData reads the quiz data from a CSV file.
func processQuizData(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseQuizData(file)
}

// parseQuizData parses the CSV file and returns the quiz data.
func parseQuizData(file *os.File) (map[string]string, error) {
	quiz := make(map[string]string)

	r := csv.NewReader(file)
	quizData, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range quizData {
		if len(record) < 2 {
			log.Printf("Invalid record (less than two fields): %v", record)
			continue
		}
		q := strings.ToLower(strings.TrimSpace(record[0]))
		a := strings.ToLower(strings.TrimSpace(record[1]))
		quiz[q] = a
	}

	return quiz, nil
}

func main() {
	quizData, err := processQuizData("questions.csv")
	if err != nil {
		log.Fatalf("Error processing quiz data: %v", err)
	}

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("You have 30 seconds to complete the quiz.")
	fmt.Println("Press ENTER to start.")
	fmt.Scanln()

	// Auotmatically creates a new goroutine to run the quiz
	// When the time expires, a signal is sent to the channel
	timer := time.NewTimer(30 * time.Second)

	quiz := Quiz{Questions: quizData}
	quiz.run(timer)
}
