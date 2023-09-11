package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans  string
}

var print = fmt.Print

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func openfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Ohhhh nooooo!!!Error in opening in file please try again after sometime!!!"))
	}
	reader := csv.NewReader(file)
	parsefile(reader)
}
func parseline(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func parsefile(filename *csv.Reader) {
	lines, err := filename.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Ooops!!!Some unexpected parsing error..."))
	}
	questions := parseline(lines)
	displayquiz(questions)
}

func displayquiz(questions []problem) {
	timelimit := 10 * len(questions)
	timer := time.NewTimer(time.Duration(timelimit) * time.Second)
	correct := 0
problemloop:
	for i, prob := range questions {
		fmt.Printf("%d: %s :: \n", i+1, prob.ques)
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerch:
			if strings.ToLower(answer) == prob.ans {
				correct++
			}
		}
	}
	evaluate(correct, len(questions))
}

func evaluate(score int, total int) {
	percentage := float64(score) / float64(total)
	if percentage > 0.8 {
		fmt.Println(fmt.Sprintf("Great man you nailed it you cracked it by answering %d question correct out of %d", score, total))
	} else if percentage > 0.6 && percentage < 0.8 {
		fmt.Println(fmt.Sprintf("Hey good job man you nearly reached by answering %d questions correct out of %d...", score, total))
	} else {
		fmt.Println(fmt.Sprintf("Oops...only %d questions answered correctly out of %d you need to improve your skills", score, total))
	}
}

func main() {
	print("Welcome to the quiz contest!!!\nThis quiz will test your knowledge and lets dive into the contest!!!")
	print("Select the topic you want to attend the quiz\n1. Verbal\n2. Math\n3. General Knowledge\nJust enter the option number eg: 1 for Verbal ----")
	var choice int
	fmt.Scanf("%d\n", &choice)
	switch choice {
	case 1:
		openfile("verbal.csv")
	case 2:
		openfile("math.csv")
	case 3:
		openfile("gk.csv")
	default:
		print("Since you have entered invalid option we will assign a random topic for test!!!")
		openfile("verbal.csv")
	}
}
