package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var selectionPoints = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

func checkOutcome(a string, b string) int {
	switch a {
	case "A":
		break
	case "B":
		break
	case "C":
		break

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(filePath string) {
	readFile, err := os.Open(filePath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	var calories []int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line != "" {
			cal, err := strconv.Atoi(line)
			check(err)
			sum += cal
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}
	calories = append(calories, sum)

	readFile.Close()

	return calories
}

func main() {
	fmt.Println("")
}
