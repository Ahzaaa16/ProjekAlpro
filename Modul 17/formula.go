package main

import "fmt"

func main() {
	var n int
	var i int
	var suku float64
	var jumlah float64

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	jumlah = 0

	for i = 1; i <= n; i++ {
		if i%2 == 1 {
			suku = 1.0 / float64(2*i-1)
		} else {
			suku = -1.0 / float64(2*i-1)
		}

		jumlah = jumlah + suku
	}

	fmt.Printf("Hasil PI: %.7f\n", 4*jumlah)

	i = 1
	jumlah = 0

	for {
		if i%2 == 1 {
			suku = 1.0 / float64(2*i-1)
		} else {
			suku = -1.0 / float64(2*i-1)
		}

		jumlah = jumlah + suku

		if (1.0 / float64(2*i+1)) <= 0.00001 {
			break
		}

		i = i + 1
	}

	fmt.Printf("Hasil PI: %.10f\n", 4*jumlah)
	fmt.Println("Pada i ke:", i)
}
