package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var banyakTopping int
	var i int
	var x, y float64
	var toppingPizza int
	var pi float64

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&banyakTopping)

	toppingPizza = 0

	for i = 1; i <= banyakTopping; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			toppingPizza = toppingPizza + 1
		}
	}

	pi = 4.0 * float64(toppingPizza) / float64(banyakTopping)

	fmt.Println("Topping pada Pizza:", toppingPizza)
	fmt.Printf("PI : %.10f\n", pi)
}
