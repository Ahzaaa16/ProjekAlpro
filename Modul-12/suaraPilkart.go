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

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	var i int
	for i = 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}
