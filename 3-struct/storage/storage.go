package storage

import (
	"demo/json/bins"
	"encoding/json"
	"fmt"
	"os"
)

type Storage interface {
	Save(binsList bins.BinList)
	Read()
}

type LocalStorage struct {
	filename string
}

func NewLocalStorage(filename string) *LocalStorage {
	return &LocalStorage{
		filename: filename,
	}
}

func (storage *LocalStorage) Save(binsList bins.BinList) {
	data, err := json.Marshal(binsList)
	if err != nil {
		fmt.Println("Не удалось распарсить файл: ", err)
		return
	}

	err = os.WriteFile(storage.filename, data, 0644)

	if err != nil {
		fmt.Println("Не удалось записать файл :", err)
		return
	}
}

func (storage *LocalStorage) Read() {
	data, err := os.ReadFile(storage.filename)
	if err != nil {
		fmt.Println("Не удалось прочитать файл: ", err)
		return
	}

	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println("Не удалось получить данные из файла: ", err)
		return
	}

	fmt.Println(binList)
}
