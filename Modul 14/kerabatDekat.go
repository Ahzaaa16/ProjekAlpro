package main

import "fmt"

func selectionSortAsc(arr []int) {
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

		var data []int
		data = make([]int, m)

		for j = 0; j < m; j++ {
			fmt.Scan(&data[j])
		}

		selectionSortAsc(data)

		var pertama bool
		pertama = true

		for j = 0; j < m; j++ {
			if data[j]%2 != 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(data[j])
				pertama = false
			}
		}

		for j = m - 1; j >= 0; j-- {
			if data[j]%2 == 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(data[j])
				pertama = false
			}
		}

		fmt.Println()
	}
}
