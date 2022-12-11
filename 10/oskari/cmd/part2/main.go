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

	for true {
		rowPos := totalCycles % 40
		totalCycles++
		if x >= rowPos-1 && x <= rowPos+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if totalCycles%40 == 0 {
			fmt.Println()
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

	readFile.Close()

}
