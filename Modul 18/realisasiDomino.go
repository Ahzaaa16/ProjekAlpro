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

func galiKartu(d *Dominoes, pembanding Domino, hasilGalian *Domino) bool {
	var ditemukan bool
	ditemukan = false

	for d.sisaKartu > 0 {
		var kartuSaatIni Domino
		kartuSaatIni = ambilKartu(d)

		var suitA1 int
		var suitB1 int
		suitA1 = gambarKartu(kartuSaatIni, 0)
		suitB1 = gambarKartu(kartuSaatIni, 1)

		var suitA2 int
		var suitB2 int
		suitA2 = gambarKartu(pembanding, 0)
		suitB2 = gambarKartu(pembanding, 1)

		if suitA1 == suitA2 || suitA1 == suitB2 || suitB1 == suitA2 || suitB1 == suitB2 {
			*hasilGalian = kartuSaatIni
			ditemukan = true
			break
		}
	}
	return ditemukan
}

func sepasangKartu(kartu1 Domino, kartu2 Domino) bool {
	var total int
	total = nilaiKartu(kartu1) + nilaiKartu(kartu2)

	var hasil bool
	if total == 12 {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func main() {
	fmt.Println(" SOAL 2 - GALIKARTU & SEPASANGKARTU")
	fmt.Println(" Disusun oleh: Rayhan Ahza Widyamukti_109082500210")

	var himpunanDomino Dominoes
	kocokKartu(&himpunanDomino)

	fmt.Println("\nKartu domino telah dikocok. Total kartu:", himpunanDomino.sisaKartu)

	var suitDicari int
	fmt.Print("\nMasukkan suit (0-6) sisi A kartu pembanding untuk digali: ")
	fmt.Scan(&suitDicari)

	var suitDicariB int
	fmt.Print("Masukkan suit (0-6) sisi B kartu pembanding untuk digali: ")
	fmt.Scan(&suitDicariB)

	var kartuPembanding Domino
	kartuPembanding.suitA = suitDicari
	kartuPembanding.suitB = suitDicariB
	kartuPembanding.nilai = suitDicari + suitDicariB
	if suitDicari == suitDicariB {
		kartuPembanding.balak = true
	} else {
		kartuPembanding.balak = false
	}

	var kartuHasilGali Domino
	var berhasil bool
	berhasil = galiKartu(&himpunanDomino, kartuPembanding, &kartuHasilGali)

	if berhasil {
		fmt.Println("\nKartu ditemukan saat menggali tumpukan:")
		fmt.Println("  Suit sisi A :", kartuHasilGali.suitA)
		fmt.Println("  Suit sisi B :", kartuHasilGali.suitB)
		fmt.Println("  Nilai kartu :", kartuHasilGali.nilai)
		fmt.Println("Sisa kartu pada tumpukan setelah digali:", himpunanDomino.sisaKartu)

		var hasilPasangan bool
		hasilPasangan = sepasangKartu(kartuPembanding, kartuHasilGali)
		fmt.Println("\nApakah kartu pembanding & kartu hasil galian berjumlah 12?", hasilPasangan)
	} else {
		fmt.Println("\nKartu dengan suit yang sama tidak ditemukan, tumpukan habis.")
	}

	fmt.Println("\nProgram selesai dijalankan oleh Rayhan Ahza Widyamukti_109082500210")
}
