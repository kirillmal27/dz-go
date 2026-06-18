package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	userWeight, userHeight := getUserInput()

	BMI := calculateBMI(userWeight, userHeight)

	fmt.Printf("Your BMI is: %.2f", BMI)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Println("Введите свой рост в метрах = ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в кг = ")
	fmt.Scan(&userWeight)

	return userWeight, userHeight
}

func calculateBMI(userWeight float64, userHeight float64) float64 {
	return userWeight / math.Pow(userHeight, 2);
}
