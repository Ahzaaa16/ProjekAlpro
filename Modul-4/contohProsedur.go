package main

import "fmt"

func main() {
	var bilangan int
	var pesan string
	fmt.Scan(&bilangan, &pesan)
	cetakPesan(pesan, bilangan)
}
func cetakPesan(M string, flag int) {
	var jenis string = ""
	if flag == 0 {
		jenis = "error"
	} else if flag == 1 {
		jenis = "warning"
	} else if flag == 2 {
		jenis = "informasi"
	}
	fmt.Println(M, jenis)
}
