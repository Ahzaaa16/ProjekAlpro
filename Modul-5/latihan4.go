package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisan(n)
}

func barisan(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	barisan(n - 1)
	fmt.Print(n, " ")
}
