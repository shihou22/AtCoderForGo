package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)
	res := hoge(x, a, b)
	fmt.Println(res)
}

func hoge(x int, a int, b int) int {
	var ans int
	ans = (x - a) % b
	return ans
}
