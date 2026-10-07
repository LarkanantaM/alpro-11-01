package main

import "fmt"

func main() {
	var p, q int
	fmt.Scan(&p, &q)

	genapAtau := p%2 == 0 || q%2 == 0
	ganjilDan := p%2 != 0 && q%2 != 0
	tidakSama := !(p == q)

	fmt.Println(genapAtau, ganjilDan, tidakSama)
}