package main

import "fmt"

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

func DaftarkanBuku(pustaka []Buku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka []Buku, n int) {
	var i, idx int

	idx = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
	}

	fmt.Println(
		pustaka[idx].judul,
		pustaka[idx].penulis,
		pustaka[idx].penerbit,
		pustaka[idx].tahun,
	)
}

func UrutBuku(pustaka []Buku, n int) {
	var i, j int
	var key Buku

	for i = 1; i < n; i++ {
		key = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka []Buku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka []Buku, n, r int) {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			fmt.Println(
				pustaka[tengah].judul,
				pustaka[tengah].penulis,
				pustaka[tengah].penerbit,
				pustaka[tengah].tahun,
				pustaka[tengah].eksemplar,
				pustaka[tengah].rating,
			)
			return
		}

		if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var n, ratingCari int

	fmt.Scan(&n)

	var pustaka []Buku
	pustaka = make([]Buku, n)

	DaftarkanBuku(pustaka, n)

	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, n)

	UrutBuku(pustaka, n)

	Cetak5Terbaru(pustaka, n)

	CariBuku(pustaka, n, ratingCari)
}
