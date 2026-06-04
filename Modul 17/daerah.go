package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	var i int
	var x, y float64
	var a, b, c, d int

	fmt.Print("Masukkan jumlah tetesan hujan: ")
	fmt.Scan(&n)

	a = 0
	b = 0
	c = 0
	d = 0

	for i = 1; i <= n; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x < 0.5 && y < 0.5 {
			a = a + 1
		} else if x >= 0.5 && y < 0.5 {
			b = b + 1
		} else if x >= 0.5 && y >= 0.5 {
			c = c + 1
		} else {
			d = d + 1
		}
	}

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", float64(a)*0.0001)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", float64(b)*0.0001)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", float64(c)*0.0001)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", float64(d)*0.0001)
}
