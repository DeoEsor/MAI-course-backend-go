package gobasics

import (
	"fmt"
	"slices"
	"sort"
)

func Test() {
	var supply Supply
	supply = Supply{
		Name:           "Gifts",
		CountOnStorage: 2,
		Price: Price{
			float32: 500.5,
			Symbol:  PriceSymbolRub,
		},
	}

	fmt.Printf("supply is name %v count on storage %v price %v\n", supply.Name, supply.CountOnStorage, supply.Price)

	var strings []string
	strings = append(strings, "")
	// a := strings[len(strings)-1]
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})
	slices.Reverse(strings)

	leaderTable := make(map[string]int)

	leaderTable["Vasya"] = 1
	leaderTable["Nikita"] = 2
	// ...
	_, ok := leaderTable["Vasysa"]
	if ok {
		println("vasyaScore exitss")
	}
	for k := range leaderTable {
		delete(leaderTable, k)
	}
	keys := make([]string, 0, len(leaderTable))
	for k := range leaderTable {
		keys = append(keys, k)
	}
	for name, score := range leaderTable {
		fmt.Printf("tournament student name %v, score %v", name, score)
	}

	set := make(map[string]struct{})

	set["Vasysa"] = struct{}{}
	_, vasyaExists := set["Vasysa"]
	if vasyaExists {
		println("vasya exitss")
	}
}

type Price struct {
	float32
	Symbol PriceSymbol
}

type PriceSymbol string

const (
	PriceSymbolNone PriceSymbol = "none"
	PriceSymbolRub  PriceSymbol = "Rub"
)

type Supply struct {
	Name           string
	CountOnStorage int
	Price          Price
}
