package main

import (
	"advent4/internal/elftools"
	"bufio"
	"fmt"
	"os"
)

func overlaps(a *elftools.Assignment, b *elftools.Assignment) bool {
	return (a.Min <= b.Max && b.Min <= a.Max)
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		a1, a2 := elftools.LineToAssignments(line)
		if overlaps(&a1, &a2) && overlaps(&a2, &a1) {
			sum++
		}
	}

	readFile.Close()

	fmt.Println(sum)
}
