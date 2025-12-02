package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"svordy/fileAccess"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func add(a int, b int, mod int) int {
	val := (a + b) % mod
	if val < 0 {
		val += mod
	}
	return val
}

func updatePassword(dial int, offset int, dialSize int, password int) (newDial int, newPassword int) {
	newDial = add(dial, offset, dialSize)
	nonModdedDial := dial + offset
	newPassword = password
	if nonModdedDial > dialSize {
		newPassword += (nonModdedDial / dialSize)
		if nonModdedDial % dialSize == 0 {
			newPassword -= 1
		}
	}
	if nonModdedDial < 0 {
		if dial != 0 {
			newPassword ++
		}
		newPassword -= (nonModdedDial / dialSize)
		if nonModdedDial % dialSize == 0 {
			newPassword -= 1
		}
	}
	if newDial == 0 {
		newPassword++
	}
	return newDial, newPassword
}

func part1UpdatePassword(dial int, offset int, dialSize int, password int) (newDial int, newPassword int) {
	newDial = add(dial, offset, dialSize)
	newPassword = password
	if newDial == 0 {
		newPassword++
	}
	return newDial, newPassword
}

func solve(updater func(int, int, int, int) (int, int), lines []string) int {
	password := 0
	dial := 50
	dialSize := 100
	for _, val := range lines {
		if val == "" {
			continue
		} else if val[0] == 'L' {
			offset, err := strconv.Atoi(val[1:])
			check(err)
			dial, password = updater(dial, - offset, dialSize, password)
		} else if val[0] == 'R' {
			offset, err := strconv.Atoi(val[1:])
			check(err)
			dial, password = updater(dial, offset, dialSize, password)
		}
	}
	return password
}

func test() error {
	lines, err := fileAccess.FileToArrayOnNewLine("day1Test.txt")
	check(err)
	part1Password := solve(part1UpdatePassword, lines)
	if part1Password != 3 {
		return errors.New("part1Logic is broken")
	}
	part2Password := solve(updatePassword, lines)
	if part2Password != 6 {
		return errors.New("part2Logic is broken")
	}
	return nil
}


func main() {
	err := test()
	check(err)
	
	lines, err := fileAccess.FileToArrayOnNewLine("day1.txt")
	check(err)
	part1Password := solve(part1UpdatePassword, lines)
	fmt.Println("part 1:", part1Password)
	part2Password := solve(updatePassword, lines)
	fmt.Println("part 2:", part2Password)
}
