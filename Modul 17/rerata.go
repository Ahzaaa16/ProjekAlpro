package main

import "fmt"

func main() {
	var bilangan float64
	var jumlah float64
	var rataRata float64
	var banyak int

	jumlah = 0
	banyak = 0

	fmt.Scan(&bilangan)

	for bilangan != 9999 {
		jumlah = jumlah + bilangan
		banyak = banyak + 1
		fmt.Scan(&bilangan)
	}

	if banyak > 0 {
		rataRata = jumlah / float64(banyak)
		fmt.Printf("Rerata = %.2f\n", rataRata)
	} else {
		fmt.Println("Tidak ada data yang dihitung")
	}
}
