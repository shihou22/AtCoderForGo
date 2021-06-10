package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// mainProjectA()
	// mainABC086A()
	// mainABC081A()
	// mainABC081B()
	// mainABC087B()
	// mainABC083B()
	// mainABC088B()
	// mainABC085B()]
	// mainABC085C()
	// mainABC049C()
	mainABC086C()
}
func mainABC086C() {
	var n int
	fmt.Scan(&n)

	var wk [][]int
	wk = make([][]int, n)
	for index := 0; index < n; index++ {
		wk[index] = make([]int, 0)
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		wk[index] = append(wk[index], t)
		wk[index] = append(wk[index], x)
		wk[index] = append(wk[index], y)
	}
	var res = abc086C(n, wk)
	fmt.Println(res)
}

func abc086C(n int, wk [][]int) string {

	var res = "Yes"

	var time = 0
	var pX = 0
	var pY = 0
	for index := 0; index < n; index++ {
		if wk[index][0]%2 == (wk[index][1]+wk[index][2])%2 {
			time = IntAbs(wk[index][0] - time)
			pX = IntAbs(wk[index][1] - pX)
			pY = IntAbs(wk[index][2] - pY)
			if time < pX+pY {
				res = "No"
				break
			}
		} else {
			res = "No"
			break
		}
		wk[index][0] = time
		wk[index][1] = pX
		wk[index][2] = pY
	}
	return res
}

func mainABC049C() {
	var wk string
	fmt.Scan(&wk)
	var res = abc049C(wk)
	fmt.Println(res)
}

func abc049C(wk string) string {
	var res = ""

	wk = Reverse(wk)
	wk = strings.Replace(wk, "resare", "", -1)
	wk = strings.Replace(wk, "esare", "", -1)
	wk = strings.Replace(wk, "remaerd", "", -1)
	wk = strings.Replace(wk, "maerd", "", -1)

	if len(wk) == 0 {
		res = "YES"
	} else {
		res = "NO"
	}
	return res
}

//Reverse string reverse
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func mainABC085C() {
	var n, y int
	fmt.Scan(&n, &y)
	var res = abc085C(n, y)
	fmt.Println(res)
}

func abc085C(n, y int) string {
	var res = ""

	for tenThous := 0; tenThous <= n; tenThous++ {
		for fiveThous := 0; fiveThous <= n; fiveThous++ {
			if tenThous+fiveThous > n {
				break
			}
			var thous = n - (tenThous + fiveThous)
			if tenThous+fiveThous+thous == n && tenThous*10000+fiveThous*5000+thous*1000 == y {
				res = strconv.Itoa(tenThous) + " " + strconv.Itoa(fiveThous) + " " + strconv.Itoa(thous)
				return res
			}

		}
	}

	return "-1 -1 -1"
}

func mainABC085B() {

	var n int
	fmt.Scan(&n)

	var wk []int
	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in)
		wk = append(wk, in)
	}
	var res = abc085B(n, wk)
	fmt.Println(res)
}

func abc085B(n int, wk []int) int {

	var res = 0

	sort.Sort(sort.Reverse(sort.IntSlice(wk)))
	var tmp = wk[0]
	res++
	for i := 0; i < n; i++ {
		if wk[i] != tmp {
			tmp = wk[i]
			res++
		}
	}

	return res
}

func mainABC088B() {
	var n int
	fmt.Scan(&n)
	var wk []int
	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in)
		wk = append(wk, in)
	}
	var res = abc088B(n, wk)
	fmt.Println(res)
}

func abc088B(n int, wk []int) int {
	var res = 0

	/*
		sort.Sort()にReverseを渡す
		sort.Reverse()にIntSliceを渡す
	*/
	sort.Sort(sort.Reverse(sort.IntSlice(wk)))

	var alice = 0
	var bob = 0
	for i := 0; i < len(wk); i++ {
		if i%2 == 0 {
			alice += wk[i]
		} else {
			bob += wk[i]
		}
	}

	res = alice - bob
	return res
}

func mainABC083B() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var res = abc083B(n, a, b)
	fmt.Println(res)
}

func abc083B(n, a, b int) int {

	var res = 0
	for i := 1; i <= n; i++ {

		var wk = i
		var total = 0
		for wk > 0 {
			total += wk % 10
			wk /= 10
		}
		if total >= a && total <= b {
			res += i
		}

	}
	return res
}

func mainABC087B() {
	var a, b, c, x int

	fmt.Scan(&a, &b, &c, &x)

	var res = abc087B(a, b, c, x)
	fmt.Println(res)
}

func abc087B(a, b, c, x int) int {
	var res = 0
	for hund5 := 0; hund5 <= a; hund5++ {
		for hund1 := 0; hund1 <= b; hund1++ {
			for five := 0; five <= c; five++ {
				if hund5*500+hund1*100+five*50 == x {
					res++
				}
			}

		}
	}
	return res
}

func mainABC081B() {
	var n int
	fmt.Scan(&n)

	var wk []int

	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in)
		wk = append(wk, in)
	}

	var res = abc081B(n, wk)
	fmt.Println(res)

}

func abc081B(n int, wk []int) int {

	var condition = true
	var res = 0
	for condition {
		for i := 0; i < len(wk); i++ {
			if wk[i]%2 == 0 {
				wk[i] /= 2
			} else {
				condition = false
				break
			}
		}
		if condition {
			res++
		}
	}
	return res
}

func mainABC081A() {
	var strS string
	fmt.Scan(&strS)

	var res = abc081A(strS)
	fmt.Println(res)
}

func abc081A(strS string) int {
	var res int
	for i := 0; i < len(strS); i++ {
		if string([]rune(strS)[i:i+1]) == "1" {
			res++
		}
	}
	return res
}

func mainABC086A() {

	var a, b int
	fmt.Scan(&a, &b)
	var res = abc086A(a, b)
	fmt.Println(res)
}

func abc086A(a, b int) string {
	var res string
	if a*b%2 == 0 {
		res = "Even"
	} else {
		res = "Odd"
	}
	return res
}

func mainProjectA() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var strS string
	fmt.Scan(&strS)
	var res = projectA(a, b, c, strS)
	fmt.Println(res)
}

func projectA(a, b, c int, strS string) string {
	var intWk = a + b + c
	var res = strconv.Itoa(intWk) + " " + strS
	return res
}

//SampleStruct privateで構造体宣言
type SampleStruct struct {
	Index  int
	Height int
}

/*
sort用
GO 1.6ではsort#Sliceが使えないのでこちらで対応
*/
type samples []SampleStruct

func (u samples) Len() int {
	return len(u)
}

func (u samples) Less(i, j int) bool {
	return u[i].Height > u[j].Height
}

func (u samples) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
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
