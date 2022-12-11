package main

import (
	"bufio"
	"fmt"
	"os"
)

func isVisible(grid [][]int, x int, y int) bool {
	treeHeight := grid[x][y]

	// check to top
	for i := x - 1; i >= 0; i-- {
		if grid[i][y] >= treeHeight {
			break
		} else if i == 0 {
			return true
		}
	}
	// check to left
	for i := y - 1; i >= 0; i-- {
		if grid[x][i] >= treeHeight {
			break
		} else if i == 0 {
			return true
		}
	}
	// check to bottom
	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] >= treeHeight {
			break
		} else if i == len(grid)-1 {
			return true
		}
	}
	// check to right
	for i := y + 1; i < len(grid[y]); i++ {
		if grid[x][i] >= treeHeight {
			break
		} else if i == len(grid)-1 {
			return true
		}
	}
	return false
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
	amountOfVisibleTrees := (len(grid)+len(grid[0]))*2 - 4
	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]
		for j := 1; j < len(row)-1; j++ {
			if isVisible(grid, i, j) {
				amountOfVisibleTrees++
			}
		}
	}
	fmt.Println(amountOfVisibleTrees)
}
