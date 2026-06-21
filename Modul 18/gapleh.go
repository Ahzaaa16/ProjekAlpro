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

// ---------- Operasi dasar mesin domino ----------

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

func cariKartuSambung(tangan []Domino, ujungKiri int, ujungKanan int) (int, string, bool, bool) {
	var indeksTerpilih int
	var sisiTerpilih string
	var perluBalik bool
	var ditemukan bool

	indeksTerpilih = -1
	sisiTerpilih = ""
	perluBalik = false
	ditemukan = false

	var i int
	for i = 0; i < len(tangan); i++ {
		var sA int
		var sB int
		sA = gambarKartu(tangan[i], 0)
		sB = gambarKartu(tangan[i], 1)

		if sA == ujungKanan {
			indeksTerpilih = i
			sisiTerpilih = "kanan"
			perluBalik = false
			ditemukan = true
			break
		}
		if sB == ujungKanan {
			indeksTerpilih = i
			sisiTerpilih = "kanan"
			perluBalik = true
			ditemukan = true
			break
		}
		if sB == ujungKiri {
			indeksTerpilih = i
			sisiTerpilih = "kiri"
			perluBalik = false
			ditemukan = true
			break
		}
		if sA == ujungKiri {
			indeksTerpilih = i
			sisiTerpilih = "kiri"
			perluBalik = true
			ditemukan = true
			break
		}
	}
	return indeksTerpilih, sisiTerpilih, perluBalik, ditemukan
}

// hapusDariTangan menghapus kartu pada indeks tertentu dari slice tangan.
func hapusDariTangan(tangan []Domino, indeks int) []Domino {
	var tanganBaru []Domino
	var i int
	for i = 0; i < len(tangan); i++ {
		if i != indeks {
			tanganBaru = append(tanganBaru, tangan[i])
		}
	}
	return tanganBaru
}

func main() {
	fmt.Println(" SOAL 3 - IMPLEMENTASI PERMAINAN GAPLEH")
	fmt.Println(" Disusun oleh: Rayhan Ahza Widyamukti_109082500210")

	var himpunanDomino Dominoes
	kocokKartu(&himpunanDomino)

	var jumlahKartuPerPemain int
	fmt.Print("\nMasukkan jumlah kartu yang dibagikan untuk setiap pemain (misal 5): ")
	fmt.Scan(&jumlahKartuPerPemain)

	var tanganP1 []Domino
	var tanganP2 []Domino

	var i int
	for i = 0; i < jumlahKartuPerPemain; i++ {
		tanganP1 = append(tanganP1, ambilKartu(&himpunanDomino))
	}
	for i = 0; i < jumlahKartuPerPemain; i++ {
		tanganP2 = append(tanganP2, ambilKartu(&himpunanDomino))
	}

	fmt.Println("\nKartu Pemain 1 (P1):")
	for i = 0; i < len(tanganP1); i++ {
		fmt.Println("  [", i, "] (", tanganP1[i].suitA, ",", tanganP1[i].suitB, ") nilai:", tanganP1[i].nilai)
	}
	fmt.Println("Kartu Pemain 2 (P2):")
	for i = 0; i < len(tanganP2); i++ {
		fmt.Println("  [", i, "] (", tanganP2[i].suitA, ",", tanganP2[i].suitB, ") nilai:", tanganP2[i].nilai)
	}
	fmt.Println("Sisa kartu di tumpukan stok:", himpunanDomino.sisaKartu)

	var indeksPembuka int
	fmt.Print("\nMasukkan indeks kartu P1 yang akan dipakai untuk membuka rangkaian: ")
	fmt.Scan(&indeksPembuka)

	var rangkaian []Domino
	rangkaian = append(rangkaian, tanganP1[indeksPembuka])
	tanganP1 = hapusDariTangan(tanganP1, indeksPembuka)

	var ujungKiri int
	var ujungKanan int
	ujungKiri = gambarKartu(rangkaian[0], 0)
	ujungKanan = gambarKartu(rangkaian[0], 1)

	fmt.Println("\nRangkaian dibuka dengan kartu (", ujungKiri, ",", ujungKanan, ")")

	var giliran int
	giliran = 2 // setelah P1 membuka, giliran berikutnya P2

	var pemainSkip int
	pemainSkip = 0

	for len(tanganP1) > 0 && len(tanganP2) > 0 {
		var tanganAktif []Domino
		if giliran == 1 {
			tanganAktif = tanganP1
		} else {
			tanganAktif = tanganP2
		}

		var idx int
		var sisi string
		var balik bool
		var dapat bool
		idx, sisi, balik, dapat = cariKartuSambung(tanganAktif, ujungKiri, ujungKanan)

		if dapat {
			var kartuSambung Domino
			kartuSambung = tanganAktif[idx]

			if sisi == "kanan" {
				if balik {
					var temp int
					temp = kartuSambung.suitA
					kartuSambung.suitA = kartuSambung.suitB
					kartuSambung.suitB = temp
				}
				rangkaian = append(rangkaian, kartuSambung)
				ujungKanan = kartuSambung.suitB
			} else {
				if balik {
					var temp int
					temp = kartuSambung.suitA
					kartuSambung.suitA = kartuSambung.suitB
					kartuSambung.suitB = temp
				}
				var rangkaianBaru []Domino
				rangkaianBaru = append(rangkaianBaru, kartuSambung)
				rangkaianBaru = append(rangkaianBaru, rangkaian...)
				rangkaian = rangkaianBaru
				ujungKiri = kartuSambung.suitA
			}

			if giliran == 1 {
				tanganP1 = hapusDariTangan(tanganP1, idx)
			} else {
				tanganP2 = hapusDariTangan(tanganP2, idx)
			}

			fmt.Println("\nPemain", giliran, "menyambung kartu (", kartuSambung.suitA, ",", kartuSambung.suitB, ") di sisi", sisi)
			fmt.Println("Rangkaian sekarang: ujung kiri =", ujungKiri, ", ujung kanan =", ujungKanan)
			pemainSkip = 0
		} else {
			fmt.Println("\nPemain", giliran, "tidak punya kartu yang cocok, lewati giliran.")
			pemainSkip = pemainSkip + 1
			if pemainSkip >= 2 {
				fmt.Println("\nKedua pemain sama-sama tidak bisa menyambung. Permainan berhenti (buntu).")
				break
			}
		}

		if giliran == 1 {
			giliran = 2
		} else {
			giliran = 1
		}
	}

	if len(tanganP1) == 0 {
		fmt.Println("\nPemain 1 (P1) menghabiskan seluruh kartunya. P1 MENANG!")
	} else if len(tanganP2) == 0 {
		fmt.Println("\nPemain 2 (P2) menghabiskan seluruh kartunya. P2 MENANG!")
	} else {
		var totalP1 int
		var totalP2 int
		for i = 0; i < len(tanganP1); i++ {
			totalP1 = totalP1 + nilaiKartu(tanganP1[i])
		}
		for i = 0; i < len(tanganP2); i++ {
			totalP2 = totalP2 + nilaiKartu(tanganP2[i])
		}
		fmt.Println("\nSisa nilai kartu P1:", totalP1)
		fmt.Println("Sisa nilai kartu P2:", totalP2)
		if totalP1 < totalP2 {
			fmt.Println("Permainan buntu, P1 MENANG (sisa nilai kartu lebih kecil)!")
		} else if totalP2 < totalP1 {
			fmt.Println("Permainan buntu, P2 MENANG (sisa nilai kartu lebih kecil)!")
		} else {
			fmt.Println("Permainan buntu dan SERI!")
		}
	}

	fmt.Println("\nProgram selesai dijalankan oleh Rayhan Ahza Widyamukti_109082500210")
}
