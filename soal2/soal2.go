package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(getPassword("abcd", "lowq"))
}

func getPassword(firstPass, level string) string {
	if len(firstPass) < 3 {
		return ("firstPass harus lebih dari 3")
	} else {
		var arrlow = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		var arrmed = [28]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		var arrstr = [10]string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}
		var newPass = firstPass
		if level == "low" {
			var i = len(newPass)
			for i < 8 {
				var index = randInt(0, 9)
				newPass = newPass + arrlow[index]
				i = len(newPass)
			}
			return (firstPass + " " + newPass)
		} else if level == "medium" {
			var i = len(newPass)
			var index [4]string
			for i < 8 {
				index[0] = arrlow[randInt(0, 9)]
				index[1] = arrmed[randInt(0, 28)]
				index[2] = arrlow[randInt(0, 9)]
				index[3] = arrmed[randInt(0, 28)]
				var y = randInt(0, 4)
				newPass = newPass + index[y]
				i = len(newPass)
			}
			return (firstPass + " " + newPass)
		} else if level == "strong" {
			newPass = newPass + arrstr[randInt(0, 10)]
			var i = len(newPass)
			var index [3]string
			for i < 8 {
				index[0] = arrlow[randInt(0, 9)]
				index[1] = arrmed[randInt(0, 10)]
				index[2] = arrstr[randInt(0, 10)]
				var y = randInt(0, 3)
				newPass = newPass + index[y]
				i = len(newPass)
			}
			return (firstPass + " " + newPass)
		} else {
			return ("salah level")
		}
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
