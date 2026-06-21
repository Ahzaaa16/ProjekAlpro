// Rayhan Ahza Widyamukti_109082500210
package main

import "fmt"

type Domino struct {
	suitA int
	suitB int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu     [28]Domino 
	sisaKartu int       
}

func kocokKartu(d *Dominoes) {
	var indeks int
	indeks = 0
	var i int
	for i = 0; i <= 6; i++ {
		var j int
		for j = i; j <= 6; j++ {
			d.kartu[indeks].suitA = i
			d.kartu[indeks].suitB = j
			d.kartu[indeks].nilai = i + j
			if i == j {
				d.kartu[indeks].balak = true
			} else {
				d.kartu[indeks].balak = false
			}
			indeks = indeks + 1
		}
	}
	d.sisaKartu = 28

	var kunciAcak int
	kunciAcak = 5
	var k int
	for k = 0; k < 28; k++ {
		var tukar int
		tukar = (k * kunciAcak) % 28
		var temp Domino
		temp = d.kartu[k]
		d.kartu[k] = d.kartu[tukar]
		d.kartu[tukar] = temp
	}
}

func ambilKartu(d *Dominoes) Domino {
	var kartuTerambil Domino
	if d.sisaKartu > 0 {
		kartuTerambil = d.kartu[d.sisaKartu-1]
		d.sisaKartu = d.sisaKartu - 1
	}
	return kartuTerambil
}

func gambarKartu(kartu Domino, suit int) int {
	var hasil int
	if suit == 0 {
		hasil = kartu.suitA
	} else {
		hasil = kartu.suitB
	}
	return hasil
}

func nilaiKartu(kartu Domino) int {
	var hasil int
	hasil = kartu.nilai
	return hasil
}

func main() {
	fmt.Println(" SOAL 1 - MESIN ABSTRAK DOMINO")
	fmt.Println(" Disusun oleh: Rayhan Ahza Widyamukti_109082500210")

	var himpunanDomino Dominoes
	kocokKartu(&himpunanDomino)

	fmt.Println("\nKartu domino telah dikocok. Total kartu:", himpunanDomino.sisaKartu)

	var jumlahAmbil int
	fmt.Print("\nMasukkan jumlah kartu yang ingin diambil dari tumpukan: ")
	fmt.Scan(&jumlahAmbil)

	var i int
	for i = 1; i <= jumlahAmbil; i++ {
		if himpunanDomino.sisaKartu <= 0 {
			fmt.Println("Kartu pada tumpukan sudah habis.")
			break
		}
		var kartuAmbil Domino
		kartuAmbil = ambilKartu(&himpunanDomino)

		var suitSisiA int
		var suitSisiB int
		suitSisiA = gambarKartu(kartuAmbil, 0)
		suitSisiB = gambarKartu(kartuAmbil, 1)

		var totalNilai int
		totalNilai = nilaiKartu(kartuAmbil)

		fmt.Println("\nKartu ke-", i, "berhasil diambil:")
		fmt.Println("  Suit sisi A :", suitSisiA)
		fmt.Println("  Suit sisi B :", suitSisiB)
		fmt.Println("  Nilai kartu :", totalNilai)
		fmt.Println("  Balak?      :", kartuAmbil.balak)
		fmt.Println("  Sisa kartu pada tumpukan:", himpunanDomino.sisaKartu)
	}

	fmt.Println("\nProgram selesai dijalankan oleh Rayhan Ahza Widyamukti_109082500210")
}
