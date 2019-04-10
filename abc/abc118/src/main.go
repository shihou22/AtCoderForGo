package main

import (
	"fmt"
	"math"
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
		// var n int
		// fmt.Scan(&n)
		// a := make([]float64, n)
		// b := make([]string, n)
		// for i := 0; i < n; i++ {
		// 	fmt.Scan(&a[i])
		// 	fmt.Scan(&b[i])
		// }
		// res := solveB(a, b)
		// fmt.Println(res)
	}

	//solveC
	{

		fmt.Scan(&nC)
		fmt.Scan(&aC)
		fmt.Scan(&bC)
		fmt.Scan(&cC)
		wk := make([]int, nC)
		for i := 0; i < nC; i++ {
			fmt.Scan(&wk[i])
		}
		res := solveC(wk)
		fmt.Println(int(res))
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

var nC, aC, bC, cC int

func solveC(wk []int) float64 {

	return dfsC(0, 0, 0, 0, wk)

}

func dfsC(currentI int, a int, b int, c int, wk []int) float64 {

	if nC == currentI {
		if a > 0 && b > 0 && c > 0 {
			return float64(math.Abs(float64(aC-a) + math.Abs(float64(bC-b)+math.Abs(float64(cC-c)-30))))
		}
		return 99999999
	}
	val0 := dfsC(currentI+1, a, b, c, wk)
	val1 := dfsC(currentI+1, a+wk[currentI], b, c, wk) + 10
	val2 := dfsC(currentI+1, a, b+wk[currentI], c, wk) + 10
	val3 := dfsC(currentI+1, a, b, c+wk[currentI], wk) + 10

	return math.Min(math.Min(val0, val1), math.Min(val2, val3))
}
