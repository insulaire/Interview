package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci2(10))
}

func fibonacci(n int) []int {
	res := []int{0, 1}
	for i := 2; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

func fibonacci2(n int) int {
	a, b := 0, 1
	for i := 2; i < n; i++ {
		t := b
		b = a + b
		a = t
	}
	return b + a
}
