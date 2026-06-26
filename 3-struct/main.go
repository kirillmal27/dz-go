package main

import (
	"demo/json/bins"
	"demo/json/storage"
	"time"
)

func main() {
	binList := bins.NewBindList()

	bin := bins.NewBin("1", true, time.Now(), "123123")
	binList.Bins = append(binList.Bins, *bin)
	storage.SaveBins("bins.json", *binList)

	bin2 := bins.NewBin("2", true, time.Now(), "qweqweqwe")
	binList.Bins = append(binList.Bins, *bin2)

	storage.SaveBins("bins.json", *binList)
}
