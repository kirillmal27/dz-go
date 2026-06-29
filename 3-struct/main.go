package main

import (
	"demo/json/bins"
	"demo/json/storage"
	"time"
)

func main() {
	binList := bins.NewBindList()
	var localStorage storage.Storage = storage.NewLocalStorage("bins.json")

	bin := bins.NewBin("1", true, time.Now(), "123123")
	binList.Bins = append(binList.Bins, *bin)
	localStorage.Save(*binList)

	bin2 := bins.NewBin("2", true, time.Now(), "qweqweqwe")
	binList.Bins = append(binList.Bins, *bin2)

	localStorage.Save(*binList)
}
