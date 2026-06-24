package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (a *account) generatePassword(length int) {
	res := make([]rune, length)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	a.password = string(res)
}

func (a account) outputPssword(acc *account) {
	fmt.Println(acc)
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)

	if len(login) == 0 {
		return nil, errors.New("invalid login")
	}

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if len(acc.password) == 0 {
		acc.generatePassword(12)
	}

	return acc, nil
}

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")

	account1, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(account1)
}

var letterRunes = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789-!*")

func promtData(promt string) string {
	fmt.Print(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
