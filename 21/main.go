package main

import "fmt"

func Fibonacci(n int) int {
	if n < 0 || n > 39 {
		panic("params is error")
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		c := a + b
		a, b = b, c
	}
	return a
}

func main() {
	fmt.Println(Fibonacci2(1))
	fmt.Println(Fibonacci2(2))
	fmt.Println(Fibonacci2(3))
	fmt.Println(Fibonacci2(4))
	fmt.Println(Fibonacci2(5))
	fmt.Println(Fibonacci2(6))
	fmt.Println(Fibonacci2(7))
	fmt.Println(Fibonacci2(39))
}
