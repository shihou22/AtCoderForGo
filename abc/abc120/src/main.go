package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// mainA()
	mainB()
}

func mainA() {
	var a, b, c int64

	fmt.Scan(&a, &b, &c)
	var res = solveA(a, b, c)
	fmt.Print(res)

}

func solveA(a int64, b int64, c int64) int64 {

	var res = b / a

	if res > c {
		return c
	}
	return res
}

func mainB() {

	var a, b, k int
	fmt.Scan(&a, &b, &k)
	var res = solveB(a, b, k)
	fmt.Println(res)
}

func solveB(a int, b int, k int) int64 {
	var i = 1
	var res = int64(0)
	var resSlice []int
	var loop = int(math.Min(float64(a), float64(b)))
	for i <= loop {
		if a%i == 0 && b%i == 0 {
			resSlice = append(resSlice, i)
		}
		i++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(resSlice)))
	res = int64(resSlice[k-1])
	return res
}
