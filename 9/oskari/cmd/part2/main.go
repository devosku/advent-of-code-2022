package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	Direction string
	Amount    int
}

// Clamp returns f clamped to [low, high]
func clamp(f, low, high int) int {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

func getNewTailPosition(tailPos [2]int, headPos [2]int) [2]int {

	// if touching
	if headPos[0] >= tailPos[0]-1 &&
		headPos[0] <= tailPos[0]+1 &&
		headPos[1] >= tailPos[1]-1 &&
		headPos[1] <= tailPos[1]+1 {
		return tailPos
	}
	newTailPos := tailPos

	newTailPos[0] += clamp(headPos[0]-tailPos[0], -1, 1)
	newTailPos[1] += clamp(headPos[1]-tailPos[1], -1, 1)

	return newTailPos

}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var moves []move

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		direction, _ := strconv.Atoi(parts[1])
		moves = append(moves, move{Direction: parts[0], Amount: direction})
	}

	readFile.Close()

	knots := make([][2]int, 10)
	for i := range knots {
		knots[i] = [2]int{0, 0}
	}
	visitedPositions := make(map[[2]int]bool, 0)
	visitedPositions[knots[9]] = true

	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			switch move.Direction {
			case "L":
				knots[0][0]--
			case "R":
				knots[0][0]++
			case "U":
				knots[0][1]--
			case "D":
				knots[0][1]++
			}
			for i := 1; i < len(knots); i++ {
				pos := getNewTailPosition(knots[i], knots[i-1])
				knots[i] = pos
			}
			visitedPositions[knots[len(knots)-1]] = true
		}
	}

	fmt.Println(len(visitedPositions))
}
