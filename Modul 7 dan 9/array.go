package main

import (
	"fmt"
)

func akar(x float64) float64 {
	var tebakan float64 = x
	var i int
	for i = 0; i < 10; i++ {
		tebakan = 0.5 * (tebakan + x/tebakan)
	}
	return tebakan
}

func main() {
	var n int
	fmt.Scan(&n)

	var arr [100]int
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i = 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	for i = 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Scan(&hapus)
	for i = hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n = n - 1

	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var jumlah int = 0
	for i = 0; i < n; i++ {
		jumlah = jumlah + arr[i]
	}
	var rata float64 = float64(jumlah) / float64(n)
	fmt.Println(rata)

	var total float64 = 0
	for i = 0; i < n; i++ {
		var selisih float64 = float64(arr[i]) - rata
		total = total + selisih*selisih
	}
	var sd float64 = akar(total / float64(n))
	fmt.Println(sd)

	var cari int
	var frek int = 0
	fmt.Scan(&cari)
	for i = 0; i < n; i++ {
		if arr[i] == cari {
			frek = frek + 1
		}
	}
	fmt.Println(frek)
}
