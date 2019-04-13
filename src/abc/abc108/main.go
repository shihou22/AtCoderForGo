package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	// mainA()
	// mainB()
	mainC()
}

func mainA() {
	var n int
	fmt.Scan(&n)
	var res = solveA(n)
	fmt.Println(res)
}

func solveA(n int) int {

	var odd = 0
	var even = 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return even * odd
}

func mainB() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	var res = solveB(x1, y1, x2, y2)
	fmt.Println(res)
}

func solveB(x1, y1, x2, y2 int) string {
	var x3 = x2 - (y2 - y1)
	var y3 = y2 + (x2 - x1)
	var x4 = x3 - (y3 - y2)
	var y4 = y3 + (x3 - x2)
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(x3))
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(y3))
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(x4))
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(y4))
	return buffer.String()

}

func mainC() {
	var n, k int
	fmt.Scan(&n, &k)
	var res = solveC(n, k)
	fmt.Println(res)
}

func solveCX(n, k int) int {

	/*
		num[x] = kで割ってxあまる数が1以上N以下に何個あるか
		mod k　することによって、N以下の整数をグルーピングする
	*/
	var num = make([]int, k)
	for i := 1; i <= n; i++ {
		//1-nまでのi mod k を全てカウント
		num[i%k]++
	}

	var res = 0
	/*
		mod k の結果を0 - (k-1)まで
		k | a+b
		を前提にすると、
		 aがmod kのリザルトなので、
		 b,c共にmod kのリザルトとして何を選べばよいのか？
		 と固定できる
	*/
	for a := 0; a < k; a++ {
		var b = (k - a) % k
		var c = (k - a) % k
		if (b+c)%k == 0 {
			/*
				a,b,c, mod k のリザルトの個数(1-nまでの間の整数をグルーピングした結果なので)
			*/
			res += num[a] * num[b] * num[c]
		}
	}
	return res
}

func solveC(n, k int) int {

	var num = make([]int, k)

	for i := 1; i <= n; i++ {
		num[i%k]++
	}

	var res = 0
	for a := 0; a < k; a++ {
		/*
		* (b+a) mod k = 0
		* b mod k = -a mod k
		* b mod k + k mod k = k mod k + -a mod k
		* 9 mod 4 = 1 = 5 mod k
		 */
		// var b = (k - a) % k
		// var c = (k - a) % k
		var b = k - a
		var c = k - a
		if b >= k || c >= k {
			b = b % k
			c = c % k
		}
		if (b+c)%k != 0 {
			continue
		}
		res += num[a] * num[b] * num[c]
	}
	return res

}

/*
* TLE-version
 */
func solveC2(n, k int) int {

	var res = make(map[string]int)

	for i := 1; i <= n; i++ {
		var a = i

		var bL = make([]int, 0)
		var cL = make([]int, 0)

		var sabun = 0
		if a%k == 0 {
			sabun = k
		} else {
			sabun = k - a%k
		}
		for j := sabun; j <= n; j += k {
			bL = append(bL, j)
			cL = append(cL, j)
		}

		for j := 0; j < len(bL); j++ {

			for j2 := 0; j2 < len(cL); j2++ {
				if (bL[j]+cL[j2])%k == 0 {
					res[strconv.Itoa(a)+":"+strconv.Itoa(bL[j])+":"+strconv.Itoa(cL[j2])] = 1
				}
			}

		}
	}

	return len(res)

}
