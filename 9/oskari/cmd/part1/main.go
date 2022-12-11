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

	headPosition := [2]int{0, 0}
	tailPosition := [2]int{0, 0}

	visitedPositions := make(map[[2]int]bool, 0)
	visitedPositions[tailPosition] = true

	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			switch move.Direction {
			case "L":
				headPosition[0]--
			case "R":
				headPosition[0]++
			case "U":
				headPosition[1]--
			case "D":
				headPosition[1]++
			}
			tailPosition = getNewTailPosition(tailPosition, headPosition)
			visitedPositions[tailPosition] = true
		}
	}

	fmt.Println(len(visitedPositions))
}
