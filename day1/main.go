package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// RunSolution - https://adventofcode.com/2020/day/1
// Find the two entries in 'expense-report.txt' that sum to 2020
// The answer is those two entries multiplied together
func RunSolution() error {
	expenseReportBytes, err := ioutil.ReadFile("./day1/expense-report.txt")
	if err != nil {
		return err
	}

	expenseItems, err := getExpenseLineItems(string(expenseReportBytes))
	if err != nil {
		return err
	}

	result, err := find2020Expense(expenseItems)

	if err != nil {
		return err
	}

	fmt.Printf("Day 1 Result: %d", result)

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

func find2020Expense(expenseItems []int) (int, error) {
	for idx1, expense1 := range expenseItems {
		for idx2, expense2 := range expenseItems {
			if idx1 != idx2 && expense1+expense2 == 2020 {
				return expense1 * expense2, nil
			}
		}
	}

	return 0, fmt.Errorf("Couldn't find two numbers in the expense report adding to 2020")
}
