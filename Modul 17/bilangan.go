package main

import "fmt"

func main() {
	var x, data string
	var n, i int
	var ditemukan bool
	var posisi int
	var jumlah int

	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah data (n): ")
	fmt.Scan(&n)

	ditemukan = false
	posisi = -1
	jumlah = 0

	for i = 1; i <= n; i++ {
		fmt.Print("Data ke-", i, ": ")
		fmt.Scan(&data)

		if data == x {
			jumlah = jumlah + 1
			ditemukan = true

			if posisi == -1 {
				posisi = i
			}
		}
	}

	fmt.Println()
	fmt.Println("Jawaban:")

	// a
	if ditemukan {
		fmt.Println("a. String x ada dalam kumpulan data.")
	} else {
		fmt.Println("a. String x tidak ada dalam kumpulan data.")
	}

	// b
	if posisi != -1 {
		fmt.Println("b. String x pertama kali ditemukan pada posisi", posisi)
	} else {
		fmt.Println("b. String x tidak ditemukan.")
	}

	// c
	fmt.Println("c. Jumlah string x dalam kumpulan data =", jumlah)

	// d
	if jumlah >= 2 {
		fmt.Println("d. Ya, terdapat sedikitnya dua string x.")
	} else {
		fmt.Println("d. Tidak, terdapat kurang dari dua string x.")
	}
}
