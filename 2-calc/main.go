package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	calculate("MED", "1, 2, 3,4, 5, 6, 9, 1, 2")
}

func calculate(operation string, numbers string) {
	sliceNumbers := strings.Split(numbers, ",")

	fmt.Println(sliceNumbers)
	
	var result int

	switch operation {
		case "AVG":
			result = calculateAvg(sliceNumbers)
		case "SUM":
			result = calculateSum(sliceNumbers)
		case "MED":
			result = calculateMed(sliceNumbers)
		default:
			fmt.Println("Не правильная операция")
	}

	fmt.Println(result)
}

func calculateAvg(numbers []string) int {
	return calculateSum(numbers) / len(numbers)
}

func calculateSum(numbers []string) int {
	var result int
	for _, value := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			continue
		}
		result += num
	}

	return result
}
func calculateMed(numbers []string) int {
	var ints []int

	for _, value := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			continue
		}
		ints = append(ints, num)
	}

	if len(ints) == 0 {
		return 0
	}

	length := len(ints)
	
	slices.Sort(ints)

	if length%2 != 0 {
		return ints[length/2]
	}

	return (ints[length/2-1] + ints[length/2]) / 2
}