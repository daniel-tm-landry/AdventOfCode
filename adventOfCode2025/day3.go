package main

import (
	"errors"
	"fmt"
	"log"
	"math"
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

func test(checker func(string) int,testFile string, expected int) error {
	input := getInput(testFile)
	result := solve(checker, input)
	if result != expected {
		return errors.New(fmt.Sprintf("expected %d, result was %d", expected, result))
	}
	return nil
}

func solve(checker func(string) int, input []string) int {
	result := 0
	for _, val := range input {
		result += checker(val)
	}
	return result
}

func part1Checker(line string) int {
	return generalChecker(line, 2)
}

func part2Checker(line string) int {
	return generalChecker(line, 12)
}

func generalChecker(line string, numDigits int) int {
	digits := make([]int, numDigits)
	inds := make([]int, numDigits)
	for i, _ := range inds {
		startInd := 0
		if i != 0 {
			startInd = inds[i - 1] + 1
		}
		for j, val := range line[startInd:len(line) - len(inds) + i + 1] {
			digit := int(val - '0')
			if digit > digits[i] {
				inds[i] = j + startInd
				digits[i] = digit
			}
		}
	}
	result := 0
	for i, val := range digits {
		result += val * int(math.Pow10(len(digits) - i - 1))
	}

	return result
}

func main() {
	var err error
	err = test(part1Checker, "day3Test1.txt", 98)
	check(err)
	err = test(part1Checker, "day3Test2.txt", 357)
	check(err)
	err = test(part2Checker, "day3Test3.txt", 811111111119)
	check(err)
	err = test(part2Checker, "day3Test1.txt", 987654321111)
	check(err)
	err = test(part2Checker, "day3Test2.txt", 3121910778619)
	check(err)
	
	input := getInput("day3.txt")
	result := solve(part2Checker, input)
	fmt.Println(result)
}
