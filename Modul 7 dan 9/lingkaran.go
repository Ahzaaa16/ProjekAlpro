package main

import (
	"fmt"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, l Lingkaran) bool {
	var dx int = t.x - l.pusat.x
	var dy int = t.y - l.pusat.y
	var jarakKuadrat int = dx*dx + dy*dy
	var radiusKuadrat int = l.r * l.r

	if jarakKuadrat <= radiusKuadrat {
		return true
	}
	return false
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	var diL1 bool = dalamLingkaran(t, l1)
	var diL2 bool = dalamLingkaran(t, l2)

	if diL1 && diL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}