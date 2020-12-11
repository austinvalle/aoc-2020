package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// RunSolution - https://adventofcode.com/2020/day/1
func RunSolution() error {
	expenseReportBytes, err := ioutil.ReadFile("./day1/expense-report.txt")
	if err != nil {
		return err
	}

	expenseItems, err := getExpenseLineItems(string(expenseReportBytes))
	if err != nil {
		return err
	}

	result1, err := findPart1Solution(expenseItems)
	if err != nil {
		return err
	}
	fmt.Printf("Day 1, Part 1 Result: %d\n", result1)

	result2, err := findPart2Solution(expenseItems)
	if err != nil {
		return err
	}
	fmt.Printf("Day 1, Part 2 Result: %d\n", result2)

	return nil
}

func getExpenseLineItems(expenseReport string) (expenseLineItems []int, err error) {
	lines := strings.Split(expenseReport, "\r\n")
	expenseLineItems = make([]int, 0, len(lines))

	for _, lineItem := range lines {
		if len(lineItem) == 0 {
			continue
		}
		amount, err := strconv.Atoi(lineItem)
		if err != nil {
			return nil, err
		}
		expenseLineItems = append(expenseLineItems, amount)
	}

	return expenseLineItems, nil
}

// O(n) solution
func findPart1Solution(expenseItems []int) (int, error) {
	visitedExpenses := make(map[int]bool)
	for _, expense := range expenseItems {
		matchingExpense := 2020 - expense
		if _, ok := visitedExpenses[matchingExpense]; ok {
			return expense * matchingExpense, nil
		}

		visitedExpenses[expense] = true
	}

	return 0, fmt.Errorf("Couldn't find two numbers in the expense report adding to 2020")
}

// O(n^3) solution, I'm lazy
func findPart2Solution(expenseItems []int) (int, error) {
	for idx1, expense1 := range expenseItems {
		for idx2, expense2 := range expenseItems[idx1:] {
			for _, expense3 := range expenseItems[idx2:] {
				if expense1+expense2+expense3 == 2020 {
					return expense1 * expense2 * expense3, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("Couldn't find three numbers in the expense report adding to 2020")
}
