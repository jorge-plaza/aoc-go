package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Move struct {
	quantity int
	from     int
	to       int
}

//go:embed example.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, "\r", "")
}

func main() {
	stacks, _ := parseInput(input)
	for _, val := range stacks {
		fmt.Printf("-> %s\n", val)
	}
}

func parseInput(input string) (stacks [][]string, moves []Move) {
	split := strings.Split(input, "\n\n")

	state := split[0]
	var oversized [][]string
	for _, row := range strings.Split(state, "\n") {
		oversized = append(oversized, strings.Split(row, ""))
	}
	oRows, oCols := len(oversized), len(oversized[0])

	for c := 0; c < oCols-1; c++ {
		if oversized[oRows-1][c] != " " {
			// hit a column with values... move up from here
			var stack []string
			for r := oRows - 2; r >= 0; r-- {
				char := oversized[r][c]
				if char != " " {
					stack = append(stack, char)
				}
			}
			stacks = append(stacks, stack)
		}
	}
	for _, move := range strings.Split(split[1], "\n") {
		step := Move{}
		_, err := fmt.Sscanf(string(move), "move %d from %d to %d", &step.quantity, &step.from, &step.to)
		if err != nil {
			panic(err)
		}
		//maybe -1 to be 0 indexed
		moves = append(moves, step)
	}
	for _, line := range strings.Split(input, "\n") {
		stacks = append(stacks, strings.Split(line, ","))
	}
	return stacks, moves
}
