//Rayhan Ahza Widyamukti_109082500210

package main

import "fmt"

func main() {
	var suara [21]int
	var input int
	var totalMasuk int
	var suaraSah int

	for {
		fmt.Scan(&input)

		totalMasuk++

		if input == 0 {
			break
		}

		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
	}

	var ketua int
	var wakil int
	var max1 int
	var max2 int
	var i int

	for i = 1; i <= 20; i++ {

		if suara[i] > max1 {
			max2 = max1
			wakil = ketua

			max1 = suara[i]
			ketua = i

		} else if suara[i] > max2 {
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
