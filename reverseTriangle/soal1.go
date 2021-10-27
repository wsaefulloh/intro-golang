package main

import (
	"fmt"
)

func cetakSegitiga(nilai int) {
	var value = ((nilai * 2) - 2) / 2
	var space = ""
	var star = ""
	for i := value; i > 0; i = i - 1 {
		var kolom = value - i
		space = ""
		star = ""
		for j := 0; j < kolom; j++ {
			space = space + " "
		}
		for y := 0; y < i; y++ {
			star = star + "*"
		}
		fmt.Println(space + star + "*" + star)
	}
	fmt.Println(space + " *")
}

func main() {
	cetakSegitiga(8)
}
