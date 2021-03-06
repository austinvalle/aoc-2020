package main

import (
	"flag"
	"fmt"

	"github.com/austinvalle/aoc-2020/day1"
	"github.com/austinvalle/aoc-2020/day2"
	"github.com/austinvalle/aoc-2020/day3"
	"github.com/austinvalle/aoc-2020/day4"
)

func main() {
	dayNumberPtr := flag.Int("d", 1, "Number of day to run.")
	flag.Parse()
	fmt.Printf("Running day #%d solution\n", *dayNumberPtr)

	var err error
	switch *dayNumberPtr {
	case 1:
		err = day1.RunSolution()
	case 2:
		err = day2.RunSolution()
	case 3:
		err = day3.RunSolution()
	case 4:
		err = day4.RunSolution()
	}

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
