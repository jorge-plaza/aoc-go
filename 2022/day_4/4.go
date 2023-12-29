package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, "\r", "")
}

func main() {
	parsedInput1 := parseInput1(input)
	result := useCase1(parsedInput1)
	fmt.Printf("Use case 1 -> %d\n", result)
	result2 := useCase2(parsedInput1)
	fmt.Printf("Use case 2 -> %d\n", result2)
}

func parseInput1(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ","))
	}
	return ans
}

func useCase1(parsedInput [][]string) int {
	var result int = 0
	for _, group := range parsedInput {
		firstElve := group[0]
		secondElve := group[1]
		sectionsFirst := [2]int{}
		sectionsSecond := [2]int{}
		for i, number := range strings.Split(firstElve, "-") {
			sectionsFirst[i], _ = strconv.Atoi(number)
		}
		for i, number := range strings.Split(secondElve, "-") {
			sectionsSecond[i], _ = strconv.Atoi(number)
		}
		if overlaps(sectionsFirst, sectionsSecond) {
			result++
		}

	}
	return result
}

func useCase2(parsedInput [][]string) int {
	var result int = 0
	fullOverlaps := 0
	for _, group := range parsedInput {
		firstElve := group[0]
		secondElve := group[1]
		sectionsFirst := [2]int{}
		sectionsSecond := [2]int{}
		for i, number := range strings.Split(firstElve, "-") {
			sectionsFirst[i], _ = strconv.Atoi(number)
		}
		for i, number := range strings.Split(secondElve, "-") {
			sectionsSecond[i], _ = strconv.Atoi(number)
		}
		if overlaps(sectionsFirst, sectionsSecond) {
			fullOverlaps++
		}
		if meh(sectionsFirst, sectionsSecond) {
			result++
		}

	}
	return result
}

func overlaps(first [2]int, second [2]int) bool {
	if first[0] <= second[0] && first[1] >= second[1] {
		return true
	} else if second[0] <= first[0] && second[1] >= first[1] {
		return true
	} else {
		return false
	}
}

func meh(first, second [2]int) bool {
	if first[0] > second[0] {
		first, second = second, first
	}
	return first[1] >= second[0]
}
