package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// mainA()
	// mainB()
	mainC()
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

func mainB() {

	var a, b, c int64
	fmt.Scan(&a, &b, &c)
	var res = solveB(a, b, c)

	fmt.Println(res)
}

func solveB(a, b, c int64) int64 {
	const ConstNum = 1000000007
	return a * b % ConstNum * c % ConstNum
}

//privateで構造体宣言
type StudentC struct {
	Index  int
	Height int
}

/*
sort用
GO 1.6ではsort#Sliceが使えないのでこちらで対応
*/
type students []StudentC

func (u students) Len() int {
	return len(u)
}

func (u students) Less(i, j int) bool {
	return u[i].Height > u[j].Height
}

func (u students) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func mainC() {
	var n int
	var wk []StudentC
	fmt.Scan(&n)

	wk = make([]StudentC, n)

	for i := 0; i < n; i++ {
		var b int
		fmt.Scan(&b)
		wk[i].Index = i + 1
		wk[i].Height = b
		// wk[i] = StudentC[Index:i+1, Height:b]
	}

	var res = solveC(n, wk)
	for i := 0; i < len(wk); i++ {
		fmt.Println(res[i])
	}
}

/*
テストのために、fmt.Printlnで直接出力するのではなく
[]stringに詰めてreturnしている
*/
func solveC(n int, wk []StudentC) []string {
	/*
		構造体のsliceをsortする
	*/
	// sort.SliceStable(wk, func(i, j int) bool {
	// 	return wk[i].Height > wk[j].Height
	// })

	studentsList := students(wk)
	sort.Sort(studentsList)

	res := make([]string, 0)
	for i := 0; i < len(wk); i++ {
		/*
			strconvで数値を文字列へ
		*/
		res = append(res, strconv.Itoa(wk[i].Index))
	}
	return res
}
