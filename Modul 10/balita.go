package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	var i int

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i = 1; i < n; i = i + 1 {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]	
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var i int
	var total float64

	total = 0.0

	for i = 0; i < n; i = i + 1 {
		total = total + arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var i int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i = 0; i < n; i = i + 1 {
		fmt.Println("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata = rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}