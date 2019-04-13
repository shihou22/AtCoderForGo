package main

import (
	"fmt"
)

func main() {
	// mainA()
	mainB()
}

func mainA() {
	var numH, numW, h, w int
	fmt.Scan(&numH, &numW, &h, &w)
	var res = solveA(numH, numW, h, w)
	fmt.Print(res)
}

func solveA(numH int, numW int, h int, w int) int {

	var res = 0
	res = (numH - h) * (numW - w)
	return res
}

func mainBInputTest(bArrayArg []int, aArrayArg [][]int, n int, m int, c int) int {

	var bArray []int
	for i := 0; i < m; i++ {
		bArray = append(bArray, bArrayArg[i])
	}
	//多次元配列使う前にはmake
	aArray := make([][]int, n)
	for outer := 0; outer < n; outer++ {
		/*
		* 長さ0で作成
		* これで要素の後ろに追加できるようになる
		 */
		aArray[outer] = make([]int, 0)
		for i := 0; i < m; i++ {
			// aArray[outer] = append(aArray[outer], aArrayArg[outer][i])
			aArray[outer] = append(aArray[outer], i)
		}
	}

	var res = solveB(bArray, aArray, n, m, c)

	return res
}

func mainB() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	var bArray []int
	for i := 0; i < m; i++ {
		var b int
		fmt.Scan(&b)
		bArray = append(bArray, b)
	}
	//多次元配列使う前にはmake
	aArray := make([][]int, n)
	for outer := 0; outer < n; outer++ {
		/*
		* 長さ0で作成
		* これで要素の後ろに追加できるようになる
		 */
		aArray[outer] = make([]int, 0)
		for i := 0; i < m; i++ {
			var b int
			fmt.Scan(&b)
			// aArray[outer][i] = b
			aArray[outer] = append(aArray[outer], b)

		}
	}

	var res = solveB(bArray, aArray, n, m, c)
	fmt.Println(res)
}

func solveB(bArray []int, aArray [][]int, n int, m int, c int) int {
	var res, point = 0, 0
	for outer := 0; outer < n; outer++ {
		point = 0
		for i := 0; i < m; i++ {
			point += aArray[outer][i] * bArray[i]
		}
		if point+c > 0 {
			res++
		}
	}
	return res
}
