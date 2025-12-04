package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"svordy/fileAccess"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getInput(testFile string) [][]string {
	input, err := fileAccess.FileToArrayOnNewLine(testFile)
	check(err)
	input = input[:len(input)-1]
	result := make([][]string, len(input), len(input[0]))
	for i, val := range input {
		result[i] = strings.Split(val, "")
	}
	return result
}

func test(checker func([][]string) int,testFile string, expected int) error {
	input := getInput(testFile)
	result := solve(checker, input)
	if result != expected {
		return errors.New(fmt.Sprintf("expected %d, result was %d", expected, result))
	}
	return nil
}

func solve(checker func([][]string) int, input [][]string) int {
	result := 0
	result += checker(input)
	return result
}

func getNeighbors(i int, j int, arr [][]string) int {
	neighbors := 0
	for x := i - 1; x <= i + 1; x++ {
		if x < 0 || x >= len(arr) {
			continue
		}
		for y := j - 1; y <= j + 1; y++ {
			if y < 0 || y >= len(arr[x]) || (x == i && y == j) {
				continue
			}

			if arr[x][y] == "@" {
				neighbors++
			}
		}
	}
	return neighbors
}


func part1Checker(arr [][]string) int {
	total := 0
	for i, line := range arr {
		for j, val := range line {
			if val == "@" {
				neighbors := getNeighbors(i, j, arr)
				if neighbors < 4 {
					total ++
				}
			}
		}
	}
	return total
}

func checkAndRemove(arr [][]string) (int, [][]string) {
	total := 0
	nextIter := make([][]string, len(arr), len(arr[0]))
	for i, line := range nextIter {
		line = make([]string, len(arr[0]))
		nextIter[i] = line
		for j, val := range line {
			val = arr[i][j]
			nextIter[i][j] = val
		}
	}
	for i, line := range arr {
		for j, val := range line {
			if val == "@" {
				neighbors := getNeighbors(i, j, arr)
				if neighbors < 4 {
					total ++
					nextIter[i][j] = "."
				}
			}
		}
	}
	return total, nextIter
}

func part2Checker(arr [][]string) int {
	total := 0
	nextIter := arr
	toAdd := 0
	for {
		toAdd, nextIter = checkAndRemove(nextIter)
		total += toAdd
		if toAdd < 1 {
			break
		}
	}
	return total
}

func main() {
	var err error
	err = test(part1Checker, "day4Test1.txt", 13)
	check(err)
	err = test(part2Checker, "day4Test1.txt", 43)
	check(err)
	
	input := getInput("day4.txt")
	result := solve(part2Checker, input)
	fmt.Println(result)
}
