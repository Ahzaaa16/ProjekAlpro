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

func main() {
	var data []int
	var bilangan int

	data = make([]int, 0)

	for {
		fmt.Scan(&bilangan)

		if bilangan < 0 {
			break
		}

		data = append(data, bilangan)
	}

	insertionSort(data)

	var i int

	for i = 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	if len(data) <= 2 {
		fmt.Println("Data berjarak tetap")
		return
	}

	var jarak int
	var tetap bool

	jarak = data[1] - data[0]
	tetap = true

	for i = 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
