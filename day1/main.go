package day1

import (
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

	_, err = getExpenseLineItems(string(expenseReportBytes))
	if err != nil {
		return err
	}

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
