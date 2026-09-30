package main

import "fmt"

func main() {
	var r float64
	var luas float64
	pi := 3.14

	fmt.Scan(&r)
	luas = pi * r * r
	fmt.Println(luas)
}