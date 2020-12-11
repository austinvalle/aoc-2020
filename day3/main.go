package day3

import (
	"bufio"
	"fmt"
	"io"
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
	mapFile, err := os.Open("./day3/map.txt")
	if err != nil {
		return err
	}

	// Part 1
	numOfTrees := calculateTreesInRoute(mapFile, 3, 1)
	fmt.Printf("Day 3, Part 1 Result: %d trees encountered\n", numOfTrees)

	// Part 2
	trees1 := calculateTreesInRoute(mapFile, 1, 1)
	trees2 := calculateTreesInRoute(mapFile, 3, 1)
	trees3 := calculateTreesInRoute(mapFile, 5, 1)
	trees4 := calculateTreesInRoute(mapFile, 7, 1)
	trees5 := calculateTreesInRoute(mapFile, 1, 2)
	part2Result := trees1 * trees2 * trees3 * trees4 * trees5
	fmt.Printf("Day 3, Part 2 Result: %d\n", part2Result)

	return nil
}

func calculateTreesInRoute(mapFile *os.File, spacesRight int, spacesDown int) int {
	// Set file pointer to beginning to allow reuse of file
	mapFile.Seek(0, io.SeekStart)

	fileScanner := bufio.NewScanner(mapFile)
	// Skip the first line, since the toboggon starts by moving down
	fileScanner.Scan()

	// Starts at top left (using zero since we'll use as index)
	xCor := 0
	numOfTrees := 0

	// Will attempt to skip down the amount of spaces, if no more lines will return false
	for scanLines(fileScanner, spacesDown) {
		mapLine := fileScanner.Text()
		if len(mapLine) < 1 {
			break
		}
		xCor += spacesRight

		// This will account for the arboreal genetics and biome stability (┛ಠ_ಠ)┛彡┻━┻
		if dups := xCor / len(mapLine); dups > 0 {
			mapLine = expandMapLine(mapLine, dups)
		}

		nextPosCharacter := mapLine[xCor : xCor+1]

		if nextPosCharacter == "#" {
			numOfTrees++
		}
	}

	return numOfTrees
}

func scanLines(scanner *bufio.Scanner, num int) bool {
	for i := 0; i < num; i++ {
		if !scanner.Scan() {
			return false
		}
	}
	return true
}

func expandMapLine(mapPattern string, dups int) string {
	var sb strings.Builder

	for i := 0; i <= dups; i++ {
		sb.WriteString(mapPattern)
	}
	return sb.String()
}
