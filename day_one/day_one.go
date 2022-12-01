package day_one

import (
	"sort"
	"strconv"
	"strings"
)

func CountCalories(input string) []int {
	result := make([]int, 0)
	split := strings.Split(input, "\n")
	total := 0
	for _, line := range split {
		if line == "" {
			if total == 0 {
				continue
			}
			result = append(result, total)
			total = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		total += val
	}
	if total != 0 {
		result = append(result, total)
	}

	return result
}

func FindHighestCalorie(input []int) int {
	if len(input) == 0 {
		return -1
	}
	sort.Ints(input)
	return input[len(input)-1]
}

func SumTop3Calories(input []int) int {
	sort.Ints(input)
	switch len(input) {
	case 0:
		return -1
	case 1:
		return input[0]
	case 2:
		return input[0] + input[1]
	default:
		return input[len(input)-1] + input[len(input)-2] + input[len(input)-3]
	}
}
