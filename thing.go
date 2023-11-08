package main

import (
	"math/rand"
	"strings"
)

type Thing struct {
	SomeString  string    `json:"some_string"`
	SomeInt     int       `json:"some_int"`
	SomeFloat64 float64   `json:"some_float64"`
	SomeSlice   []float64 `json:"some_slice"`
}

func RandomThings(src int64, n int) []Thing {

	rnd := rand.New(rand.NewSource(src)) // evenly distributed but consistent results based on src

	ret := make([]Thing, n)
	for i := 0; i < n; i++ {
		ret[i].SomeString = rndString(rnd, 1000)
		ret[i].SomeInt = rnd.Int()
		ret[i].SomeFloat64 = rnd.Float64()
		ret[i].SomeSlice = make([]float64, 0, 1000)
		for j := 0; j < 10; j++ {
			ret[i].SomeSlice = append(ret[i].SomeSlice, rnd.Float64())
		}
	}

	return ret
}

var rndLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-\"`'\\") // too lazy for multi-byte UTF8 atm 🫠

func rndString(rnd *rand.Rand, n int) string {
	var b strings.Builder
	b.Grow(n)
	c := rndLetters[rnd.Int()%len(rndLetters)]
	b.WriteRune(c) // note: won't work for chars gt 1 byte
	return b.String()
}
