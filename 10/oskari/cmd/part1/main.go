package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var instructions []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		instructions = append(instructions, line)
	}

	x := 1
	pos := 0
	currentInstruction := instructions[pos]
	parts := strings.Split(currentInstruction, " ")
	var cyclesRemaining int
	if parts[0] == "noop" {
		cyclesRemaining = 1
	} else if parts[0] == "addx" {
		cyclesRemaining = 2
	}
	totalCycles := 0
	sum := 0

	for true {
		totalCycles++
		if totalCycles == 20 || totalCycles == 60 || totalCycles == 100 || totalCycles == 140 || totalCycles == 180 || totalCycles == 220 {
			sum += totalCycles * x
		}
		cyclesRemaining--
		if cyclesRemaining <= 0 {
			parts := strings.Split(currentInstruction, " ")
			if parts[0] == "addx" {
				v, _ := strconv.Atoi(parts[1])
				x += v
			}
			pos++
			if pos >= len(instructions) {
				break
			}
			currentInstruction = instructions[pos]
			parts = strings.Split(currentInstruction, " ")
			if parts[0] == "noop" {
				cyclesRemaining = 1
			} else if parts[0] == "addx" {
				cyclesRemaining = 2
			}
		}
	}

	fmt.Println(sum)

	readFile.Close()

}
