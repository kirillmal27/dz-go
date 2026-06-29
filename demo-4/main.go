package main

import (
	"demo/password/account"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

type someFunc = func(string, string) bool

func main() {
	vault := account.NewVault()
Menu:
	for {
		userChoice := promtData([]string{
			"1: Создать аккаунт",
			"2: Найти аккаунт по URL",
			"3: Найти аккаунт по логину",
			"4: Удалить аккаунт",
			"5: Выход",
			"Выберите вариант",
		})

		switch userChoice {
		case "1":
			createAccount(vault)
		case "2":
			findAccountsByURL(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Menu
		}
	}

}

func createAccount(vault *account.Vault) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите url"})

	account1, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*account1)
}

func findAccountsByURL(vault *account.Vault) {
	url := promtData([]string{"Введите URL для поиска"})

	accounts := vault.FindAccountByURL(url)

	b, err := json.MarshalIndent(accounts, "", "    ")
	if err != nil {
		color.Red("Ошибка маршалинга: %v", err)
		return
	}
	fmt.Println(color.GreenString("%+v", string(b)))

}

func deleteAccount(vault *account.Vault) {
	url := promtData([]string{"Введите URL для удаления"})

	isDeleted := vault.DeleteAccountByURL(url)

	if isDeleted {
		color.Green("Удаление прошло успешно")
	} else {
		color.Red("Ничего не удалили")
	}
}

func promtData[T any](promt []T) string {
	var res string
	for index, elem := range promt {
		if index == len(promt)-1 {
			fmt.Printf("%s :", elem)
		}
		fmt.Println(elem)
	}
	fmt.Scanln(&res)
	return res
}
