package main

import "fmt"

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	var dalam1 bool = didalam(cx1, cy1, r1, x, y)
	var dalam2 bool = didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func didalam(cx, cy, r, x, y int) bool {
	var jarak2 int
	jarak2 = (x-cx)*(x-cx) + (y-cy)*(y-cy)

	if jarak2 <= r*r {
		return true
	}
	return false
}
