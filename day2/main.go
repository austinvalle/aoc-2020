package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type password struct {
	password        string
	requiredLetter  string
	minimumRequired int
	maximumAllowed  int
}

// RunSolution - https://adventofcode.com/2020/day/2
// Find how many passwords are valid, given the policy and password in 'passwords.txt'
func RunSolution() error {
	passwordsFile, err := os.Open("./day2/passwords.txt")
	if err != nil {
		return err
	}

	validPasswords := 0
	fileScanner := bufio.NewScanner(passwordsFile)
	for fileScanner.Scan() {
		pass, err := parsePasswordLine(fileScanner.Text())
		if err != nil {
			return err
		}

		if isValidPassword(pass) {
			validPasswords++
		}
	}

	fmt.Printf("Day 2, Part 1 Result: %d valid passwords", validPasswords)
	return nil
}

func parsePasswordLine(passwordLine string) (password, error) {
	splitStrArray := strings.SplitAfter(passwordLine, " ")

	if len(splitStrArray) < 3 {
		return password{}, fmt.Errorf("Invalid password line: %s", passwordLine)
	}
	rangeStr := strings.TrimSpace(splitStrArray[0])
	rangeArr := strings.Split(rangeStr, "-")
	if len(rangeArr) < 2 {
		return password{}, fmt.Errorf("Invalid password letter range: %s", passwordLine)
	}
	minimumRequired, err := strconv.Atoi(rangeArr[0])
	if err != nil {
		return password{}, err
	}

	maximumAllowed, err := strconv.Atoi(rangeArr[1])
	if err != nil {
		return password{}, err
	}

	return password{
		password:        strings.TrimSpace(splitStrArray[2]),
		requiredLetter:  strings.TrimSuffix(splitStrArray[1], ": "),
		minimumRequired: minimumRequired,
		maximumAllowed:  maximumAllowed,
	}, nil
}

func isValidPassword(pass password) bool {
	numOfOccurences := strings.Count(pass.password, pass.requiredLetter)
	return numOfOccurences >= pass.minimumRequired && numOfOccurences <= pass.maximumAllowed
}
