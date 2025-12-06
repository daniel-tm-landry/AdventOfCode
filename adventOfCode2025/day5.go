package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"svordy/fileAccess"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getInput(testFile string) []string {
	input, err := fileAccess.FileToArrayOnNewLine(testFile)
	check(err)
	input = input[:len(input)-1]
	return input
}

func test(checker func() bool,testFile string, expected int) error {
	input := getInput(testFile)
	result := solve(checker, input)
	if result != expected {
		return errors.New(fmt.Sprintf("expected %d, result was %d", expected, result))
	}
	return nil
}


func solve(checker func() bool, input []string) int {
	result := 0
	ranges := make([]MinMax, 1)
	i := 0
	for ; input[i] != ""; i++ {
		m := expandRange(input[i])
		shouldAdd := false
		for j, toCheck := range ranges {
			if toCheck.contains(m.low) {
				if toCheck.contains(m.high) {
					shouldAdd = false
				} else {
					ranges[j] = MinMax{toCheck.low, m.high}
					shouldAdd = false
				}
			} else if toCheck.contains(m.high) {
				ranges[j] = MinMax{m.low, toCheck.high}
				shouldAdd = false
			} else if m.contains(toCheck.low) {
				ranges[j] = m
				shouldAdd = false
			} else {
				shouldAdd = true
			}
		}
		if shouldAdd {
			ranges = append(ranges, m)
		}
	}
	ranges = ranges[1:]

	var shouldDelete []int

	for h, m := range ranges {
		killit := false
		for j, toCheck := range ranges {
			if h == j || slices.Contains(shouldDelete, j) {
				continue
			} else if toCheck.contains(m.low) {
				if toCheck.contains(m.high) {
					killit = true
				} else {
					ranges[j] = MinMax{toCheck.low, m.high}
					killit = true
				}
			} else if toCheck.contains(m.high) {
				ranges[j] = MinMax{m.low, toCheck.high}
				killit = true
			} else if m.contains(toCheck.low) {
				ranges[j] = m
				killit = true
			}
		}

		if killit {
			shouldDelete = append(shouldDelete, h)
		}
	}
	for i := len(shouldDelete) - 1; i >= 0; i-- {
		ranges = append(ranges[:shouldDelete[i]], ranges[shouldDelete[i]+1:]...)
	}
	if checker() {
		for _, set := range ranges {
			result += set.total()
		}
		return result
	} 
	
	i++
	for ; i < len(input); i++ {
		numToCheck, err := strconv.Atoi(input[i])
		check(err)
		for _, toCheck := range ranges {
			if toCheck.contains(numToCheck) {
				result++
				break
			}
		}
	}
	return result
}

type MinMax struct {
	low int
	high int
}

func (m MinMax) contains(val int) bool {
	return val >= m.low && val <= m.high
}

func (m MinMax) total() int {
	return m.high + 1 - m.low
}

func expandRange(val string) MinMax {
	bookends := strings.Split(val, "-")
	if len(bookends) != 2 {
		check(errors.New("invalid range"))
	}
	start, err := strconv.Atoi(bookends[0])
	check(err)
	end, err := strconv.Atoi(bookends[1])
	check(err)
	return MinMax{low: start, high: end}
}

func part1Checker() bool {
	return false
}

func part2Checker() bool {
	return true
}

func main() {
	var err error
	err = test(part1Checker, "day5Test1.txt", 3)
	check(err)
	err = test(part2Checker, "day5Test1.txt", 14)
	check(err)
	
	input := getInput("day5.txt")
	result := solve(part2Checker, input)
	fmt.Println(result)
}
