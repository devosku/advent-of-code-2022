package main

import (
	"advent4/internal/elftools"
	"bufio"
	"fmt"
	"os"
)

func fullyContains(a *elftools.Assignment, b *elftools.Assignment) bool {
	return (a.Min <= b.Min && a.Max >= b.Max)
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		a1, a2 := elftools.LineToAssignments(line)
		if fullyContains(&a1, &a2) || fullyContains(&a2, &a1) {
			sum++
		}
	}

	readFile.Close()

	fmt.Println(sum)
}
