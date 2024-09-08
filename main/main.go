package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a, b := 0, 1

	for i := 0; i < 10; i++ {
		for b < n {
			a, b = b, a+b
		}
		fmt.Println(b)
		a, b = b, a+b
	}
}
