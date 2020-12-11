package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type password struct {
	password       string
	requiredLetter string
	firstNumber    int
	secondNumber   int
}

// RunSolution - https://adventofcode.com/2020/day/2
// Find how many passwords are valid, given the policy and password in 'passwords.txt'
func RunSolution() error {
	passwordsFile, err := os.Open("./day2/passwords.txt")
	if err != nil {
		return err
	}

	validSledRentalPasswords := 0
	validTobogganPasswords := 0
	fileScanner := bufio.NewScanner(passwordsFile)
	for fileScanner.Scan() {
		pass, err := parsePasswordLine(fileScanner.Text())
		if err != nil {
			return err
		}

		if isValidSledRentalPassword(pass) {
			validSledRentalPasswords++
		}

		if isValidTobogganPassword(pass) {
			validTobogganPasswords++
		}
	}

	fmt.Printf("Day 2, Part 1 Result: %d valid sled rental passwords\n", validSledRentalPasswords)
	fmt.Printf("Day 2, Part 2 Result: %d valid toboggan passwords\n", validTobogganPasswords)
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
	firstNumber, err := strconv.Atoi(rangeArr[0])
	if err != nil {
		return password{}, err
	}

	secondNumber, err := strconv.Atoi(rangeArr[1])
	if err != nil {
		return password{}, err
	}

	return password{
		password:       strings.TrimSpace(splitStrArray[2]),
		requiredLetter: strings.TrimSuffix(splitStrArray[1], ": "),
		firstNumber:    firstNumber,
		secondNumber:   secondNumber,
	}, nil
}

func isValidSledRentalPassword(pass password) bool {
	numOfOccurences := strings.Count(pass.password, pass.requiredLetter)
	return numOfOccurences >= pass.firstNumber && numOfOccurences <= pass.secondNumber
}

func isValidTobogganPassword(pass password) bool {
	isInFirst := pass.password[pass.firstNumber-1:pass.firstNumber] == pass.requiredLetter
	isInSecond := pass.password[pass.secondNumber-1:pass.secondNumber] == pass.requiredLetter
	// golang doesn't support a XOR logical operator, so this is an equivalent
	return isInFirst != isInSecond
}
