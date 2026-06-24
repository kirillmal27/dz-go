package main

import "fmt"

func main() {
	transactions := []float64{}

	s := make([]int, 3, 3)
	slice2 := append(s, 4)

	fmt.Println(s)
	fmt.Println(slice2)

	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	var result float64

	for _, value := range transactions {
		result += value
	}

	fmt.Printf("Сумма всех транзакций : %.2f", result)
}

func scanTransaction() (transaction float64) {

	fmt.Println("ВВедите транзакцию")
	fmt.Scan(&transaction)
	return
}