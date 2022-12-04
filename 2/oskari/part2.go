package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var roundResultPoints = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var p2SelectionPoints = map[[2]string]int{
	{"A", "X"}: 3,
	{"A", "Y"}: 1,
	{"A", "Z"}: 2,
	{"B", "X"}: 1,
	{"B", "Y"}: 2,
	{"B", "Z"}: 3,
	{"C", "X"}: 2,
	{"C", "Y"}: 3,
	{"C", "Z"}: 1,
}

func part2() {
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		roundScore := roundResultPoints[parts[1]] + p2SelectionPoints[[2]string{parts[0], parts[1]}]
		sum += roundScore
	}

	readFile.Close()

	fmt.Println(sum)
}
