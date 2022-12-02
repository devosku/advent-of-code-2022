package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCalories(filePath string) []int {
	readFile, err := os.Open(filePath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	var calories []int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line != "" {
			cal, err := strconv.Atoi(line)
			check(err)
			sum += cal
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}
	calories = append(calories, sum)

	readFile.Close()

	return calories
}

func findMaxCalories(calories []int) int {
	max := 0
	for _, v := range calories {
		if v > max {
			max = v
		}
	}
	return max
}

func findTopThreeTotal(calories []int) int {
	sortedCalories := make([]int, len(calories))
	copy(sortedCalories, calories)
	sort.Slice(sortedCalories, func(i, j int) bool {
		return sortedCalories[i] > sortedCalories[j]
	})
	sum := 0
	for _, v := range sortedCalories[0:3] {
		sum += v
	}
	return sum
}

func main() {
	calories := readCalories("./calories.txt")
	max := findMaxCalories(calories)
	fmt.Println("Max calories:")
	fmt.Println(max)
	topThree := findTopThreeTotal(calories)
	fmt.Println("Sum of top three calories:")
	fmt.Println(topThree)
}
