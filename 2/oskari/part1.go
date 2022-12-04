package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var p1SelectionPoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var outcomePoints = map[[2]string]int{
	{"A", "X"}: 3,
	{"A", "Y"}: 6,
	{"A", "Z"}: 0,
	{"B", "X"}: 0,
	{"B", "Y"}: 3,
	{"B", "Z"}: 6,
	{"C", "X"}: 6,
	{"C", "Y"}: 0,
	{"C", "Z"}: 3,
}

func part1() {
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		roundScore := p1SelectionPoints[parts[1]] + outcomePoints[[2]string{parts[0], parts[1]}]
		sum += roundScore
	}

	readFile.Close()

	fmt.Println(sum)
}
