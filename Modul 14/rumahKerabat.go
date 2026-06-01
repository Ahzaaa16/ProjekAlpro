package main

import "fmt"

func selectionSort(arr []int) {
	var n int
	var i, j, min, temp int

	n = len(arr)

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func main() {
	var n int
	var m int
	var i, j int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		var rumah []int
		rumah = make([]int, m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for j = 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}