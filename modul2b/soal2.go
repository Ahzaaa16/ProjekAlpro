package main

import "fmt"

func main() {

	var bunga string
	var pita string
	var jumlah int
	var i int

	pita = ""
	jumlah = 0
	i = 1

	for i <= 100 {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah = jumlah + 1
		i = i + 1
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
