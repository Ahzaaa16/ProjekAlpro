package main

import (
	"fmt"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	n   int
}

func isiArray(t *tabel, n *int) {
	var i int
	var temp string

	fmt.Scan(n)
	t.n = *n

	for i = 0; i < *n; i++ {
		fmt.Scan(&temp)
		t.tab[i] = rune(temp[0])
	}
}

func cetakArray(t tabel, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print(string(t.tab[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t tabel, n int) tabel {
	var hasil tabel
	var i int

	hasil.n = n
	for i = 0; i < n; i++ {
		hasil.tab[i] = t.tab[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	balik = balikanArray(t, n)

	var i int
	for i = 0; i < n; i++ {
		if t.tab[i] != balik.tab[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks : ")
	cetakArray(tab, n)

	var hasil tabel
	hasil = balikanArray(tab, n)
	fmt.Print("Reverse teks : ")
	cetakArray(hasil, n)

	fmt.Print("Palindrom : ")
	if palindrom(tab, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
