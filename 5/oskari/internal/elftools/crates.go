package elftools

import (
	"strconv"
	"strings"
)

type CrateStack []string

// IsEmpty: check if stack is empty
func (s *CrateStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *CrateStack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *CrateStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func ParseStacks(s []string) []CrateStack {
	lastLineParts := strings.Fields(s[len(s)-1])
	amountOfCrates := len(lastLineParts)
	crates := make([]CrateStack, amountOfCrates)
	for i := len(s) - 2; i >= 0; i-- {
		line := s[i]
		offset := 3
		for j := 0; j < amountOfCrates; j++ {
			crate := string(line[j+1+(j*offset)])
			if crate == " " {
				continue
			} else {
				crates[j].Push(crate)
			}
		}
	}
	return crates
}

func ParseMoves(s []string) [][3]int {
	var moves [][3]int
	for _, line := range s {
		var move [3]int
		parts := strings.Split(line, " ")
		move[0], _ = strconv.Atoi(parts[1])
		move[1], _ = strconv.Atoi(parts[3])
		move[2], _ = strconv.Atoi(parts[5])
		moves = append(moves, move)
	}
	return moves
}
