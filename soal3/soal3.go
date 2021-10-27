package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getFilm(99))
}

func getFilm(value int) string {
	var film = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var film1 = 0
	var film2 = 0
	var result = ""
	for i, nilai1 := range film {
		for y := i + 1; y < len(film); y++ {
			var nilai2 = film[y]
			var total = nilai1 + nilai2
			if total == value {
				film1 = nilai1
				film2 = nilai2
				result = strconv.Itoa(film1) + " dan " + strconv.Itoa(film2)
			}
		}
	}
	if film1 == 0 {
		result = "tidak ada"
	}
	return (result)
}
