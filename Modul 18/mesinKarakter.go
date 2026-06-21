// Rayhan Ahza Widyamukti_109082500210
package main

import "fmt"

type MesinKarakter struct {
	teks    string
	posisi  int
	panjang int
}

func start(m *MesinKarakter, teksMasukan string) {
	m.teks = teksMasukan
	m.posisi = 0
	m.panjang = len(teksMasukan)
}

func maju(m *MesinKarakter) {
	m.posisi = m.posisi + 1
}

func eop(m *MesinKarakter) bool {
	var hasil bool
	if m.posisi >= m.panjang {
		hasil = true
	} else if m.teks[m.posisi] == '.' {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func cc(m *MesinKarakter) byte {
	var hasil byte
	hasil = m.teks[m.posisi]
	return hasil
}

func main() {
	fmt.Println(" SOAL 4 - MESIN ABSTRAK KARAKTER")
	fmt.Println(" Disusun oleh: Rayhan Ahza Widyamukti_109082500210")

	var teksMasukan string
	fmt.Print("\nMasukkan untaian karakter tanpa spasi (akhiri dengan tanda titik '.'): ")
	fmt.Scan(&teksMasukan)

	var mesin MesinKarakter

	start(&mesin, teksMasukan)
	fmt.Println("\n--- Membaca seluruh karakter ---")
	for eop(&mesin) == false {
		var karakterSaatIni byte
		karakterSaatIni = cc(&mesin)
		fmt.Printf("%c", karakterSaatIni)
		maju(&mesin)
	}
	fmt.Println()

	start(&mesin, teksMasukan)
	var jumlahKarakter int
	jumlahKarakter = 0
	for eop(&mesin) == false {
		jumlahKarakter = jumlahKarakter + 1
		maju(&mesin)
	}
	fmt.Println("\nJumlah karakter yang terbaca (tidak termasuk titik):", jumlahKarakter)

	start(&mesin, teksMasukan)
	var jumlahHurufA int
	jumlahHurufA = 0
	for eop(&mesin) == false {
		var karakterSaatIni byte
		karakterSaatIni = cc(&mesin)
		if karakterSaatIni == 'A' {
			jumlahHurufA = jumlahHurufA + 1
		}
		maju(&mesin)
	}
	fmt.Println("Jumlah huruf 'A' yang terbaca:", jumlahHurufA)

	var frekuensiA float64
	if jumlahKarakter > 0 {
		frekuensiA = float64(jumlahHurufA) / float64(jumlahKarakter)
	} else {
		frekuensiA = 0
	}
	fmt.Printf("Frekuensi kemunculan huruf 'A' terhadap seluruh karakter: %.4f\n", frekuensiA)

	start(&mesin, teksMasukan)
	var jumlahPasanganLE int
	jumlahPasanganLE = 0
	var karakterSebelumnya byte
	karakterSebelumnya = 0
	for eop(&mesin) == false {
		var karakterSaatIni byte
		karakterSaatIni = cc(&mesin)
		if karakterSebelumnya == 'L' && karakterSaatIni == 'E' {
			jumlahPasanganLE = jumlahPasanganLE + 1
		}
		karakterSebelumnya = karakterSaatIni
		maju(&mesin)
	}
	fmt.Println("Jumlah pasangan huruf 'L' diikuti 'E' (\"LE\") yang terbaca:", jumlahPasanganLE)

	fmt.Println("\nProgram selesai dijalankan oleh Rayhan Ahza Widyamukti_109082500210")
}
