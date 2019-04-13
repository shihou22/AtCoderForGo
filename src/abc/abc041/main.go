package main

import (
	"fmt"
)

func main() {

	mainA()
}

func mainA() {
	var s string
	var numI int
	fmt.Scan(&s, &numI)
	var res = solveA(s, numI)
	fmt.Println(res)
}

func solveA(s string, numI int) string {
	return s[numI-1 : numI]
}
