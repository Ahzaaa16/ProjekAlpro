package main

import "fmt"

func main() {

	var K int
	var i int
	var hasil float64
	var f float64

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	hasil = 1
	i = 0

	for i <= K {

		f = ((4*float64(i) + 2) * (4*float64(i) + 2)) /
			((4*float64(i) + 1) * (4*float64(i) + 3))

		hasil = hasil * f

		i = i + 1
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
