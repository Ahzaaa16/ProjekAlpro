package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

func isiArray(n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var i int

	for i = 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}

	return -1
}

func main() {
	var n int
	var k int
	var hasil int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil = posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
