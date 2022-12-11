package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	Items     []int
	Operation string
	Test      int
	IfTrue    int
	IfFalse   int
}

func calcWorryLevel(operation string, old int) int {
	parts := strings.Split(operation, " ")
	symbol := parts[3]
	var leftOperand int
	var rightOperand int
	if parts[2] == "old" {
		leftOperand = old
	} else {
		leftOperand, _ = strconv.Atoi(parts[2])
	}
	if parts[4] == "old" {
		rightOperand = old
	} else {
		rightOperand, _ = strconv.Atoi(parts[4])
	}
	if symbol == "+" {
		return (leftOperand + rightOperand)
	} else {
		return (leftOperand * rightOperand)
	}
}

func processMonkeyRound(monkeys []monkey, inspections []int) ([]monkey, []int) {
	newMonkeys := monkeys
	for i, monkey := range monkeys {
		for _, itemWorryLevel := range monkey.Items {
			inspections[i]++
			worryLevel := calcWorryLevel(monkey.Operation, itemWorryLevel) / 3
			var targetMonkey int
			if worryLevel%monkey.Test == 0 {
				targetMonkey = monkey.IfTrue
			} else {
				targetMonkey = monkey.IfFalse
			}
			newMonkeys[targetMonkey].Items = append(newMonkeys[targetMonkey].Items, worryLevel)
		}
		newMonkeys[i].Items = make([]int, 0)
	}
	return newMonkeys, inspections
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var monkeys []monkey

	monkeys = append(monkeys, monkey{})
	currentMonkey := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.Trim(line, " ")
		if strings.HasPrefix(line, "Starting items: ") {
			parts := strings.Split(strings.TrimPrefix(line, "Starting items: "), ", ")
			for _, p := range parts {
				item, _ := strconv.Atoi(p)
				monkeys[currentMonkey].Items = append(monkeys[currentMonkey].Items, item)
			}
		} else if strings.HasPrefix(line, "Operation: ") {
			monkeys[currentMonkey].Operation = strings.TrimPrefix(line, "Operation: ")
		} else if strings.HasPrefix(line, "Test: ") {
			parts := strings.Split(line, " ")
			test, _ := strconv.Atoi(parts[3])
			monkeys[currentMonkey].Test = test
		} else if strings.HasPrefix(line, "If true: ") {
			parts := strings.Split(line, " ")
			ifTrue, _ := strconv.Atoi(parts[5])
			monkeys[currentMonkey].IfTrue = ifTrue
		} else if strings.HasPrefix(line, "If false: ") {
			parts := strings.Split(line, " ")
			ifFalse, _ := strconv.Atoi(parts[5])
			monkeys[currentMonkey].IfFalse = ifFalse
		}

		if line == "" {
			monkeys = append(monkeys, monkey{})
			currentMonkey++
		}
	}
	readFile.Close()

	inspections := make([]int, len(monkeys))
	for i := range inspections {
		inspections[i] = 0
	}

	for i := 0; i < 20; i++ {
		monkeys, inspections = processMonkeyRound(monkeys, inspections)
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println(inspections[0] * inspections[1])
}
