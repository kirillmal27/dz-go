package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	arrayString := strings.Split(filename, ".")
	if len(arrayString) > 1 {
		if arrayString[1] != "json" {
			return nil, errors.New("Файл не json расширени")
		}

		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return data, nil
	}

	return nil, errors.New("Не корреткное название файла")
}
