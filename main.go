package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.82
	userWeight := 95.0
	BMI := userWeight / math.Pow(userHeight, 2);

	fmt.Printf("Your BMI is: %.2f", BMI)
}
