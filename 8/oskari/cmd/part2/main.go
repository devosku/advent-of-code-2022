package main

import (
	"bufio"
	"fmt"
	"os"
)

func getScenicScore(grid [][]int, x int, y int) int {
	treeHeight := grid[x][y]

	// check to top
	topDist := 0
	for i := x - 1; i >= 0; i-- {
		topDist++
		if grid[i][y] >= treeHeight {
			break
		}
	}
	// check to left
	leftDist := 0
	for i := y - 1; i >= 0; i-- {
		leftDist++
		if grid[x][i] >= treeHeight {
			break
		}
	}
	// check to bottom
	bottomDist := 0
	for i := x + 1; i < len(grid); i++ {
		bottomDist++
		if grid[i][y] >= treeHeight {
			break
		}
	}
	// check to right
	rightDist := 0
	for i := y + 1; i < len(grid[y]); i++ {
		rightDist++
		if grid[x][i] >= treeHeight {
			break
		}
	}
	return topDist * leftDist * bottomDist * rightDist
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var grid [][]int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}

	readFile.Close()
	maxScenicScore := 0
	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]
		for j := 1; j < len(row)-1; j++ {
			scenicScore := getScenicScore(grid, i, j)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	fmt.Println(maxScenicScore)
}
