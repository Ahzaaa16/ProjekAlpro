package main

import "fmt"

func insertionSort(arr []int) {
	var i, j, key int

	for i = 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func median(arr []int) int {
	var n int
	n = len(arr)

	if n%2 == 1 {
		return arr[n/2]
	}

	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	var angka int

	data = make([]int, 0)

	for {
		fmt.Scan(&angka)

		if angka == -5313 {
			break
		}

		if angka == 0 {
			var salinan []int
			salinan = make([]int, len(data))
			copy(salinan, data)

			insertionSort(salinan)
			fmt.Println(median(salinan))
		} else {
			data = append(data, angka)
		}
	}
}
