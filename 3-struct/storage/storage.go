package storage

import (
	"demo/json/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBins(filepath string, binsList bins.BinList) {
	data, err := json.Marshal(binsList)
	if err != nil {
		fmt.Println("Не удалось распарсить файл: ", err)
		return
	}

	err = os.WriteFile(filepath, data, 0644)

	if err != nil {
		fmt.Println("Не удалось записать файл :", err)
		return
	}
}

func ReadBins(filepath string) {
	data, err := os.ReadFile(filepath)
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
