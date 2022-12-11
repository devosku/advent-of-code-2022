package main

import (
	"advent7/internal/elftools"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSize(n *elftools.Node) int {
	sum := 0
	for _, file := range n.Children {
		if file.FileType == 1 {
			sum += file.Size
		} else {
			sum += getSize(file)
		}
	}
	return sum
}

func findLargerButClosestToSize(n *elftools.Node, size int, closest *int) int {
	curSize := getSize(n)
	if curSize >= size && curSize-size <= *closest {
		*closest = curSize
		for _, file := range n.Children {
			if file.FileType == 0 {
				*closest = findLargerButClosestToSize(file, size, closest)
			}
		}
	}
	return *closest
}

func main() {
	readFile, _ := os.Open("../../input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rootNode := elftools.Node{FileType: 0, Name: "root", Parent: nil, Children: make(map[string]*elftools.Node, 0)}
	var current *elftools.Node

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				if line[5:] == "/" {
					current = &rootNode
				} else if line[5:] == ".." {
					current = current.Parent
				} else {
					current = current.Children[line[5:]]
				}
			}
		} else if strings.HasPrefix(line, "dir ") {
			newDir := elftools.Node{FileType: 0, Name: line[4:], Parent: current, Children: make(map[string]*elftools.Node, 0)}
			current.Children[newDir.Name] = &newDir
		} else {
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			newFile := elftools.Node{FileType: 1, Name: parts[1], Parent: current, Size: size}
			current.Children[newFile.Name] = &newFile
		}
	}

	readFile.Close()

	currentClosest := getSize(&rootNode)
	unusedSpace := 70000000 - currentClosest
	neededSpace := 30000000 - unusedSpace
	sum := findLargerButClosestToSize(&rootNode, neededSpace, &currentClosest)
	fmt.Println(sum)
}
