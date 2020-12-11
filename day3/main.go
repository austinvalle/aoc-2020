package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type password struct {
	password       string
	requiredLetter string
	firstNumber    int
	secondNumber   int
}

// RunSolution - https://adventofcode.com/2020/day/3
func RunSolution() error {
	// How many trees would you encounter, starting at the top left, and going right 3, down 1
	mapFile, err := os.Open("./day3/map.txt")
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(mapFile)

	// Skip the first line, since the toboggon starts by moving down
	fileScanner.Scan()

	// Starts at top left (using zero since we'll use as index)
	xCor := 0

	numOfTrees := 0
	for fileScanner.Scan() {
		mapLine := fileScanner.Text()
		if len(mapLine) < 1 {
			break
		}
		xCor += 3

		// This will account for the arboreal genetics and biome stability (┛ಠ_ಠ)┛彡┻━┻
		if dups := xCor / len(mapLine); dups > 0 {
			mapLine = expandMapLine(mapLine, dups)
		}

		nextPosCharacter := mapLine[xCor : xCor+1]

		if nextPosCharacter == "#" {
			numOfTrees++
		}
	}

	fmt.Printf("Day 3, Part 1 Result: %d trees encountered", numOfTrees)
	return nil
}

func expandMapLine(mapPattern string, dups int) string {
	var sb strings.Builder

	for i := 0; i <= dups; i++ {
		sb.WriteString(mapPattern)
	}
	return sb.String()
}
