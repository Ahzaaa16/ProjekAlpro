package main

import "fmt"

type arrIkan [1000]float64

func main() {
	var data arrIkan
	var x, y int
	var i, j, index int
	var total float64
	var rata float64

	fmt.Print("Masukan jumlah ikan (x) dan jumlah wadah (y): ")
	fmt.Scan(&x, &y)

	for i = 0; i < x; i = i + 1 {
		fmt.Print("Masukan berat ikan ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	index = 0

	fmt.Println("Total berat tiap wadah:")
	for i = 0; i < y; i = i + 1 {
		total = 0.0

		for j = 0; j < x/y; j = j + 1 {
			total = total + data[index]
			index = index + 1
		}

		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	total = 0.0
	for i = 0; i < x; i = i + 1 {
		total = total + data[i]
	}

	rata = total / float64(x)

	fmt.Printf("Rata-rata berat ikan: %.2f\n", rata)
}