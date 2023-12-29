package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, "\r", "")
}

func main() {
	a := "a"
	A := "A"
	unicodea := int32(a[0])
	unicodeA := int32(A[0])
	const aPriority int32 = 1
	const APriority int32 = 27

	compartments := parseInput(input)
	useCase1(compartments, unicodeA, APriority, unicodea, aPriority)
	useCase2(parseInput2(input), unicodeA, APriority, unicodea, aPriority)
}

func useCase1(compartments [][]string, unicodeA int32, APriority int32, unicodea int32, aPriority int32) {
	var result int32 = 0
	//Use a map with auto presence check as the default value for boolean is false
	for _, compartment := range compartments {
		//fmt.Printf("[%s-%s]\n", compartment[0], compartment[1])
		found := make(map[int32]bool)
		for _, item := range compartment[0] {
			//fmt.Printf("%d-%s,", item, string(item))
			if item < 65 {
				panic("Unmanaged character")
			}
			found[item] = true
		}
		for _, item := range compartment[1] {
			if found[item] {
				//This is the repeated item
				//Capital letter
				if item >= 65 && item < 97 {
					result += item - unicodeA + APriority
				} else if item >= 97 {
					//Lower case letter
					calc := item - unicodea + aPriority
					result += calc
				}
				break
			}
		}
	}
	fmt.Printf("Turbo final result %d\n", result)
}

func useCase2(compartments []string, unicodeA int32, APriority int32, unicodea int32, aPriority int32) {
	var result int32 = 0
	for i := 0; i < len(compartments); i += 3 {
		groupSet := make(map[int]*Set)
		for j := i; j < i+3; j++ {
			groupSet[j] = NewSet()
			for _, letter := range compartments[j] {
				groupSet[j].Add(letter)
			}
		}

		int1 := groupSet[i].Intersect(groupSet[i+1])
		int2 := int1.Intersect(groupSet[i+2])
		for item := range int2.list {
			if item >= 65 && item < 97 {
				result += item - unicodeA + APriority
			} else if item >= 97 {
				//Lower case letter
				calc := item - unicodea + aPriority
				result += calc
			}
		}
	}
	fmt.Printf("Second day result %d\n", result)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]
		ans = append(ans, []string{firstHalf, secondHalf})
	}
	return ans
}

func parseInput2(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

type Set struct {
	list map[int32]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v int32) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v int32) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v int32) {
	delete(s.list, v)
}

func (s *Set) Clear() {
	s.list = make(map[int32]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int32]struct{})
	return s
}

func (s *Set) Intersect(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}
