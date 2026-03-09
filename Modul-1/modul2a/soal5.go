package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var s string

	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Scan(&s)

	fmt.Printf("%c%c%c\n", s[0]+1, s[1]+1, s[2]+1)
}
