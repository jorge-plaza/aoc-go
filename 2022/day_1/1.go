package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var m = 0

func main() {
	//number of calories of each Bicho input

	top3 := [3]int{}

	f, err := os.Open("./input1.txt")
	check(err)
	reader := bufio.NewReader(f)

	current := 0

	for true {
		line, _, err := reader.ReadLine()
		fmt.Printf("%s \n", line)
		if err == io.EOF {
			inTop(&top3, current)
			fmt.Println("current top", &top3)
			var sum int
			for _, n := range top3 {
				sum += n
			}
			println("Final sum", sum)
			panic(io.EOF)
		}
		//check(err)
		if len(line) == 0 {
			inTop(&top3, current)
			if current > m {
				m = current
			}
			current = 0
			continue
		}
		amount, err := strconv.Atoi(string(line))
		check(err)
		current += amount
	}
}

func inTop(top *[3]int, amount int) {
	fmt.Println("Elve total", amount, "current top", top)
	for i := 0; i < len(top); i++ {
		if top[i] < amount {
			top[i] = amount
			break
		}
	}

	sort.Ints(top[:])
	fmt.Println("current top", top)
}

func check(e error) {
	if e == io.EOF {
		fmt.Println(m)
	}
	if e != nil {
		panic(e)
	}
}
