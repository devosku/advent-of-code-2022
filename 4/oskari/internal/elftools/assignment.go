package elftools

import (
	"strconv"
	"strings"
)

type Assignment struct {
	Min, Max int
}

func StringToAssignment(s string) Assignment {
	parts := strings.Split(s, "-")
	partsAsInt := make([]int, len(parts))
	for i, v := range parts {
		partsAsInt[i], _ = strconv.Atoi(v)
	}
	return Assignment{Min: partsAsInt[0], Max: partsAsInt[1]}
}

func LineToAssignments(s string) (Assignment, Assignment) {
	parts := strings.Split(s, ",")
	assignment1 := StringToAssignment(parts[0])
	assignment2 := StringToAssignment(parts[1])
	return assignment1, assignment2
}
