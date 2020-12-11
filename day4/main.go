package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rules = map[string]func(string) bool{
	"byr": createYearValidator(1920, 2002),
	"iyr": createYearValidator(2010, 2020),
	"eyr": createYearValidator(2020, 2030),
	"hgt": createHeightValidator(),
	"hcl": createRegexValidator("^#(?:[0-9a-fA-F]{3}){1,2}$"),
	"ecl": createRegexValidator("^(amb|blu|brn|gry|grn|hzl|oth)$"),
	"pid": func(value string) bool {
		_, err := strconv.Atoi(value)
		return len(value) == 9 && err == nil
	},
}

// RunSolution - https://adventofcode.com/2020/day/4
func RunSolution() error {
	passportsFile, err := os.Open("./day4/passports.txt")
	if err != nil {
		return err
	}

	passportScanner := bufio.NewScanner(passportsFile)

	validPassports := 0
	currentPassport := make(map[string]string)
	for passportScanner.Scan() || len(currentPassport) > 0 {
		passportLine := passportScanner.Text()
		if len(strings.TrimSpace(passportLine)) == 0 {
			isValid := isValidPassport(currentPassport)
			if isValid {
				validPassports++
			}
			currentPassport = make(map[string]string)
			continue
		}

		pairs := strings.Split(passportLine, " ")
		for _, pair := range pairs {
			seperatorIndex := strings.Index(pair, ":")
			key := pair[:seperatorIndex]
			value := pair[seperatorIndex+1:]

			currentPassport[key] = value
		}
	}

	fmt.Printf("Day 4, Results: %d valid passports\n", validPassports)

	return nil
}

func isValidPassport(passport map[string]string) bool {
	for key, isValid := range rules {
		if !keyExists(passport, key) || !isValid(passport[key]) {
			return false
		}
	}

	return true
}

func keyExists(passport map[string]string, key string) bool {
	_, exists := passport[key]
	return exists
}

func createYearValidator(lowerYear int, upperYear int) func(string) bool {
	return func(value string) bool {
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return year >= lowerYear && year <= upperYear
	}
}

func createRegexValidator(regexPattern string) func(string) bool {
	return func(value string) bool {
		matched, _ := regexp.MatchString(regexPattern, value)
		return matched
	}
}

func createHeightValidator() func(string) bool {
	return func(value string) bool {
		if strings.Contains(value, "cm") {
			height := strings.TrimSuffix(value, "cm")
			heightNbr, err := strconv.Atoi(height)
			if err != nil || heightNbr < 150 || heightNbr > 193 {
				return false
			}
			return true
		} else if strings.Contains(value, "in") {
			height := strings.TrimSuffix(value, "in")
			heightNbr, err := strconv.Atoi(height)
			if err != nil || heightNbr < 59 || heightNbr > 76 {
				return false
			}
			return true
		}
		return false
	}
}
