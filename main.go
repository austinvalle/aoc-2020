package main

import (
	"flag"
	"fmt"

	"github.com/austinvalle/aoc-2020/day1"
)

func main() {
	dayNumberPtr := flag.Int("d", 1, "Number of day to run.")
	flag.Parse()
	fmt.Printf("Running day #%d solution\n", *dayNumberPtr)

	var err error
	switch *dayNumberPtr {
	case 1:
		err = day1.RunSolution()
	}

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}