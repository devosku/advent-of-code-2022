package main

import (
	"advent4/internal/elftools"
	"bufio"
	"fmt"
	"os"
)

func executeMoves(crates []elftools.CrateStack, moves [][3]int) []elftools.CrateStack {
	for _, move := range moves {
		var stack []string
		for i := 0; i < move[0]; i++ {
			tmp, _ := crates[move[1]-1].Pop()
			stack = append(stack, tmp)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			crates[move[2]-1].Push(stack[i])
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
