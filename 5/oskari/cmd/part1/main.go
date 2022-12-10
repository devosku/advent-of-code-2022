package main

import (
	"advent5/internal/elftools"
	"bufio"
	"fmt"
	"os"
)

func executeMoves(crates []elftools.CrateStack, moves [][3]int) []elftools.CrateStack {
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			tmp, _ := crates[move[1]-1].Pop()
			crates[move[2]-1].Push(tmp)
		}
	}
	return crates
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var moves [][3]int
	var stacks []elftools.CrateStack
	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			stacks = elftools.ParseStacks(lines)
			lines = make([]string, 0)
		} else {
			lines = append(lines, line)
		}
	}
	moves = elftools.ParseMoves(lines)

	readFile.Close()

	stacks = executeMoves(stacks, moves)

	for _, stack := range stacks {
		top, _ := stack.Pop()
		fmt.Print(top)
	}
	fmt.Println()
}
