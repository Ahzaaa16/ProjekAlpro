package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	pola(n, 1)
}

func pola(n int, i int) {
	if i > n {
		return
	} else {
		fmt.Println(bintang(i))
		pola(n, i+1)
	}
}

func bintang(j int) string {
	if j == 0 {
		return ""
	} else {
		return "*" + bintang(j-1)
	}
}
