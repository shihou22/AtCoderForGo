package main

import (
	"fmt"
)

func main() {

	// mainA()
	mainB()
	// mainC()
}

func mainA() {
	var n, x int
	fmt.Scan(&n, &x)
	var res = solveA(n, x)
	fmt.Println(res)
}

func solveA(n, x int) int {
	if n-x > x-1 {
		return x - 1
	} else {
		return n - x
	}
}

func mainB() {

	var n int
	fmt.Scan(&n)
	var res = solveB(n)
	fmt.Println(res)
}

func solveB(n int) int {

	var res = 999999999
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var wk = n - (i * j)
			if wk < 0 {
				break
			}
			var wkRes = IntAbs(i-j) + wk
			res = IntMin(res, wkRes)
		}
	}
	return res

}

func mainC() {
	var n int
	var aA []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&aA[i])
	}
	var res = solveC(n, aA)
	fmt.Println(res)
}

func solveC(n int, aA []int) int {

	var dp = make([]int, n)

	for i := 1; i < n; i++ {
		dp[i] = 1000000001
	}

	dp[0] = 0
	// dp = append(dp, 0)
	for i := 0; i < n; i++ {
		for j := 1; j <= 2; j++ {
			if i+j >= n {
				break
			}
			dp[i+j] = IntMin(dp[i]+IntAbs(aA[i+j]-aA[i]), dp[i+j])
		}
	}
	return dp[n-1]
}

//Int64Abs abstract int64
func Int64Abs(a int64) int64 {
	if a < 0 {
		a *= -1
	}
	return a
}

//IntAbs abstract int
func IntAbs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

//IntMin return minimum value
func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//IntMax return maximum value
func IntMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
