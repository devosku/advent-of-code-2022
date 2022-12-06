package part1

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findItem(line string) int {
	firstCompartment := make(map[rune]bool)
	length := len(line)

	var foundInBoth rune
	for i, char := range line {
		if i < length/2 {
			firstCompartment[char] = true
		} else if firstCompartment[char] {
			foundInBoth = char
			break
		}
	}

	if unicode.IsUpper(foundInBoth) {
		return int(foundInBoth) - 38
	} else {
		return int(foundInBoth) - 96
	}
}

func Run() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		v := findItem(line)
		sum += v
	}

	readFile.Close()

	fmt.Println(sum)
}
