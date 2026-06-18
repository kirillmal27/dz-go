package main

import "fmt"

func getUserInput() float64 {
	var input float64

	fmt.Scan(&input)

	return input
}

func convertMoney(money float64, currentCurrency string, intoCurrency string) {}

func main() {
	const usdToEur = 0.92
	const usdToRub = 91.50

	const eurToRub = usdToRub / usdToEur

	fmt.Printf("Курс USD -> EUR: %.2f\n", usdToEur)
	fmt.Printf("Курс USD -> RUB: %.2f\n", usdToRub)
	fmt.Printf("Рассчитанный курс EUR -> RUB: %.2f\n", eurToRub)
}