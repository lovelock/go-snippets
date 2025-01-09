package main

import (
	"testing"

	"github.com/brianvoe/gofakeit"
)

var pile []Item

func init() {
	pile = generateRandomItems(100000)
}

func BenchmarkSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slice_contains1(gofakeit.Name(), pile)
	}
}

func BenchmarkSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slice_contains2(gofakeit.Name(), pile)
	}
}
