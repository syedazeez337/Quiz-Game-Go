package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of question, answer")
	timeLimit := flag.Int("limit", 10, "Time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	} /*else {
		fmt.Println("Success")
	}*/
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintln("Failed to parse the provided CSV file"))
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0

	problems := parseLines(lines)
	for i, val := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, val.ques)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		
		select {
		case <-timer.C:
			fmt.Println()
			return
		case answer := <-answerCh:
			if answer == val.ans {
				counter++
			}
		}
	}

	fmt.Printf("You got %d of %d correct\n", counter, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// a new type of problems
type problem struct {
	ques string
	ans  string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
