package main

import (
	"bufio"
	"fmt"
	"os"
)

func isMarker(s string, i int) bool {
	m := make(map[rune]bool, 14)
	for _, r := range s[i : i+14] {
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	markerStart := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i := range line {
			if isMarker(line, i) {
				markerStart = i
				break
			}
		}
	}

	readFile.Close()

	fmt.Println(markerStart + 14)
}
