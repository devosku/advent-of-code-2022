package part2

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

func findCommon(a string, b string) string {
	first := make(map[rune]bool)
	common := make(map[rune]bool)

	for _, char := range a {
		first[char] = true
	}

	for _, char := range b {
		if first[char] {
			common[char] = true
		}
	}

	commonString := ""

	for key, _ := range common {
		commonString += string(key)
	}
	return commonString
}

func findItem(group [3]string) int {

	common := findCommon(group[0], group[1])
	common = findCommon(common, group[2])

	if len(common) > 1 {
		fmt.Println("Something is wrong...")
		panic(1)
	}

	var foundInAll rune
	for _, v := range common {
		foundInAll = v
	}

	if unicode.IsUpper(foundInAll) {
		return int(foundInAll) - 38
	} else {
		return int(foundInAll) - 96
	}
}

func Run() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	pos := 0

	var group [3]string

	for fileScanner.Scan() {
		pos = pos % 3
		line := fileScanner.Text()
		group[pos] = line
		if pos == 2 {
			sum += findItem(group)
		}
		pos++
	}

	readFile.Close()

	fmt.Println(sum)
}
