package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	fmt.Scan(&klubA)
	fmt.Scan(&klubB)

	var hasil [100]string
	var skorA, skorB int
	var i int = 0
	var nomor int = 1

	for {
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[i] = klubA
		} else if skorB > skorA {
			hasil[i] = klubB
		} else {
			hasil[i] = "Draw"
		}

		i = i + 1
	}

	var j int
	for j = 0; j < i; j++ {
		fmt.Println("Hasil", nomor, ":", hasil[j])
		nomor = nomor + 1
	}

	fmt.Println("Pertandingan selesai")
}
