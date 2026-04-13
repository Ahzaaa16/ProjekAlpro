package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(1, n)
}

func ganjil(i int, n int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	ganjil(i+2, n)
}
