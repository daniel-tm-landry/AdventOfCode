package main

import (
	"errors"
	"fmt"
	"log"
  "github.com/dlclark/regexp2"
	"strings"
	"strconv"
	"svordy/fileAccess"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getInput(testFile string) []string {
	input, err := fileAccess.FileToArray(testFile, ",")
	input[len(input) - 1] = input[len(input) - 1][:len(input[len(input) - 1]) - 2]
	check(err)
	return input
}

func test(testFile string, expected int, expected2 int) error {
	input := getInput(testFile)
	result := solve(part1Checker, input)
	if result != expected {
		return errors.New(fmt.Sprintf("expected %d, result was %d", expected, result))
	}
	result = solve(part2Checker, input)
	if result != expected2 {
		return errors.New(fmt.Sprintf("expected %d, result was %d", expected2, result))
	}
	return nil
}

func expandRange(val string) (string, error) {
	bookends := strings.Split(val, "-")
	result := ""
	if len(bookends) != 2 {
		return "", errors.New("invalid range")
	}
	start, err := strconv.Atoi(bookends[0])
	check(err)
	end, err := strconv.Atoi(bookends[1])
	check(err)
	for i := start; i < end; i++ {
		result += strconv.Itoa(i)
		result += ","
	}
	result += strconv.Itoa(end)
	return result, nil
}

func part1Checker(numToCheck string) int {
	re := regexp2.MustCompile(`^(\d+)\1$`, 0)
	match, err := re.MatchString(numToCheck)
	check(err)
	if match {
		result, err := strconv.Atoi(numToCheck)
		check(err)
		return result
	}
	return 0
}

func part2Checker(numToCheck string) int {
	re := regexp2.MustCompile(`^(\d+)\1+$`, 0)
	match, err := re.MatchString(numToCheck)
	check(err)
	if match {
		result, err := strconv.Atoi(numToCheck)
		check(err)
		return result
	}
	return 0
}

func checkRange(rangeToCheck string, checker func(string) int) int {
	result := 0
	for _, val := range strings.Split(rangeToCheck, ",") {
		result += checker(val)
	}
	return result
}

func solve(checker func(string) int, input []string) int {
	result := 0
	for _, val := range input {
		fullRange, err := expandRange(val)
		check(err)
		result += checkRange(fullRange, checker)
	}
	return result
}

func main() {
	var err error
	err = test("day2Test2.txt", 33, 33)
	check(err)
	err = test("day2Test.txt", 1227775554, 4174379265)
	check(err)
	
	input := getInput("day2.txt")
	result := solve(part2Checker, input)
	fmt.Println(result)
}
