package bins

import (
	"encoding/json"
	"os"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func NewBindList() *BinList {
	data, err := os.ReadFile("bins.json")
	if err != nil {
		return &BinList{
			Bins: []Bin{},
		}
	}
	var bins []Bin
	err = json.Unmarshal(data, &bins)

	return &BinList{
		Bins: bins,
	}
}
