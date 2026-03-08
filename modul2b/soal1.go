package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var berhasil bool
	var i int

	berhasil = true

	for i = 1; i <= 5; i = i + 1 {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
