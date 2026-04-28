package main

import "fmt"

type arrKelinci [1000]float64

func main() {
	var data arrKelinci
	var n int
	var i int
	var min, max float64

	fmt.Print("Masukan jumlah anak kelinci: ")
	fmt.Scan(&n)

	for i = 0; i < n; i = i + 1 {
		fmt.Print("Masukan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	min = data[0]
	max = data[0]

	for i = 1; i < n; i = i + 1 {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}