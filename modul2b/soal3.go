package main

import "fmt"

func main() {

	var kiri, kanan float64
	var selisih float64
	var oleng bool
	var total float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		total = kiri + kanan

		if total > 150 || kiri < 0 || kanan < 0 {
			break
		}

		if kiri > kanan {
			selisih = kiri - kanan
		} else {
			selisih = kanan - kiri
		}

		if selisih >= 9 {
			oleng = true
		} else {
			oleng = false
		}

		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}

	fmt.Println("Proses selesai.")
}
