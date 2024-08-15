package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func takeInputFromCSVFile() ([][]string, error) {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		return nil, fmt.Errorf("error opening csv file: %v", err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %v", err)
	}
	return lines, nil
}

func convertLinetoProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func printQuestionAndGetAnswer(problems []problem) int32 {
	correct := int32(0)
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	return correct
}

func main() {
	lines, err := takeInputFromCSVFile()
	if err != nil {
		fmt.Printf("error taking input. %v\n", err)
		os.Exit(1)
	}
	problems := convertLinetoProblems(lines)

	correctAnswers := printQuestionAndGetAnswer(problems)

	fmt.Printf("Your score %d out of %d.\n", correctAnswers, len(problems))
}
