package main

import (
	"github.com/brianvoe/gofakeit"
)

type Item struct {
	Name  string
	Age   int8
	Other string
}

func generateRandomItems(n int) []Item {
	items := make([]Item, n)

	for i := 0; i < n; i++ {
		items[i] = Item{
			Name:  gofakeit.Name(),
			Age:   int8(gofakeit.Number(0, 90)),
			Other: gofakeit.Sentence(6),
		}
	}

	return items
}

func slice_contains1(needle string, pile []Item) int {

	for i, v := range pile {
		if v.Name == needle {
			return i
		}
	}

	return 0
}

func slice_contains2(needle string, pile []Item) int {

	for i := range pile {
		if pile[i].Name == needle {
			return i
		}
	}

	return 0
}
