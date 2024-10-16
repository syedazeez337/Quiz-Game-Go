package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of question, answer")
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

	counter := 0

	problems := parseLines(lines)
	for i, val := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, val.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == val.ans {
			fmt.Println("Correct")
			counter++
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("You got %d of %d correct\n", counter, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans: line[1],
		}
	}
	return ret
}

// a new type of problems
type problem struct {
	ques string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
