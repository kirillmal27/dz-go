package main

import (
	"fmt"
	"os"
)

type bookmarkMap = map[string]string

func main() {
	bookmarksMap := bookmarkMap{}



Menu:
	for {
		userChoice := getMenu()
		switch userChoice {
			case 1:
				showMeBookmarks(bookmarksMap)
			case 2:
				bookmarksMap = inputBookmarksInfo(bookmarksMap)
			case 3:
				bookmarksMap = deleteBookmarks(bookmarksMap)
			case 4:
				break Menu
		}

			
	}
}

func getMenu() int {
	var userChoice int

	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Выход")
	fmt.Scan(&userChoice)

	return userChoice
}

func showMeBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("У вас нет закладок")
	} else {
		fmt.Println(bookmarks)
	}
}

func inputBookmarksInfo(bookmarks bookmarkMap) bookmarkMap {
	var bookmarksName string
	var bookmarkValue string

	fmt.Println("Введите название закладки: ")
	fmt.Scan(&bookmarksName)

	fmt.Println("Введите значение закладки: ")
	fmt.Scan(&bookmarkValue)

	bookmarks[bookmarksName] = bookmarkValue

	return bookmarks
}

func deleteBookmarks(bookmarks bookmarkMap) bookmarkMap {
	var bookmarksName string

	fmt.Println("Введите название закладки которую надо удалить: ")
	fmt.Scan(&bookmarksName)

	delete(bookmarks, bookmarksName)

	return bookmarks
}

func finish() {
	fmt.Println("Программа сейчас завершится с кодом 1 (ошибка).")
	os.Exit(1)
}
