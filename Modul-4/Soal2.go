package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal += 1
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int
	var maxSoal = -1
	var minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
