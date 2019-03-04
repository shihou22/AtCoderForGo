package main

import (
	"fmt"
	"time"
)

const (
	yyyymmdd = "2006/01/02"
)

func main() {
	// solveA
	{
		// var dateS, resA string
		// _, _ = fmt.Scan(&dateS)
		// resA = solveA(dateS)
		// fmt.Println(resA)
	}

	//solveB
	{
		var n int
		fmt.Scan(&n)
		a := make([]float64, n)
		b := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			fmt.Scan(&b[i])
		}
		res := solveB(a, b)
		fmt.Println(res)
	}
}

func solveA(dateS string) string {

	t1, _ := time.Parse(yyyymmdd, dateS)
	t2, _ := time.Parse(yyyymmdd, "2019/04/30")

	if t1.Unix() <= t2.Unix() {
		return "Heisei"
	} else {
		return "TBD"
	}
}

func solveB(a []float64, b []string) float64 {

	total := 0.0
	for index := 0; index < len(a); index++ {
		if b[index] == "JPY" {
			total += a[index]
		} else if b[index] == "BTC" {
			total += a[index] * 380000.0
		}
	}
	return total
}
