package main

import "fmt"

func main() {

	var berat int
	var kg int
	var sisa int
	var biayaKg int
	var biayaSisa int
	var total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000

	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total = biayaKg + biayaSisa

	if kg > 10 {
		total = biayaKg
	}

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)

}
