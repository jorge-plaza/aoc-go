package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//Total score = shape + outcome
//Shape A/X - Rock - 1
//		B/Y - Paper- 2
//		C/Z - Scissors - 3
//Outcome 	lost - 0
//			draw - 3
//			win - 6
/* E.g
A Y 2 + 6
B X 1 + 0
C Z 3 + 3
Total = 15
*/
const (
	WIN     = 6
	DRAW    = 3
	LOSE    = 0
	ROCK    = 1
	PAPER   = 2
	SCISSOR = 3
)

//go:embed input2.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, "\r", "")
}

func main() {
	plays := parseInput(input)
	result1 := useCase1(plays)
	result2 := useCase2(plays)
	fmt.Printf("Result problem 1 %d \n", result1)
	fmt.Printf("Result problem 2 %d \n", result2)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, " "))
	}
	return ans
}

func useCase1(plays [][]string) int {
	myChoices := map[string]int{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSOR,
	}
	totalScore := 0
	for _, play := range plays {
		elve := play[0]
		me := play[1]
		totalScore += myChoices[me]
		if elve == "A" {
			switch me {
			case "X":
				totalScore += DRAW
			case "Y":
				totalScore += WIN
			case "Z":
				totalScore += LOSE
			}
		} else if elve == "B" {
			switch me {
			case "X":
				totalScore += LOSE
			case "Y":
				totalScore += DRAW
			case "Z":
				totalScore += WIN
			}
		} else if elve == "C" {
			switch me {
			case "X":
				totalScore += WIN
			case "Y":
				totalScore += LOSE
			case "Z":
				totalScore += DRAW
			}
		}
	}
	return totalScore
}

/*
X you need to LOSE
Y you need to DRAW
Z you need to WIN
*/
func useCase2(plays [][]string) int {
	totalScore := 0
	myChoices := map[string]int{
		"X": LOSE,
		"Y": DRAW,
		"Z": WIN,
	}
	for _, play := range plays {
		elve := play[0]
		me := play[1]
		totalScore += myChoices[me]
		if elve == "A" {
			switch me {
			case "X":
				totalScore += SCISSOR
			case "Y":
				totalScore += ROCK
			case "Z":
				totalScore += PAPER
			}
		} else if elve == "B" {
			switch me {
			case "X":
				totalScore += ROCK
			case "Y":
				totalScore += PAPER
			case "Z":
				totalScore += SCISSOR
			}
		} else if elve == "C" {
			switch me {
			case "X":
				totalScore += PAPER
			case "Y":
				totalScore += SCISSOR
			case "Z":
				totalScore += ROCK
			}
		}
	}
	return totalScore
}
