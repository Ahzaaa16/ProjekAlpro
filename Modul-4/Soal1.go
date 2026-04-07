package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	if r == 0 {
		r = 2
	}
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	if r == 0 {
		r = 2 
	}
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, b, &p1)
	combination(a, b, &c1)
	permutation(c, d, &p2)
	combination(c, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
