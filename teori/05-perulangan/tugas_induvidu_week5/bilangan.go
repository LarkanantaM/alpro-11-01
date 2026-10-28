package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	cantik := true
	kanan := x % 10
	x = x / 10

	for x > 0 {
		kiri := x % 10
		if kiri > kanan {
			cantik = false
			break
		}
		kanan = kiri
		x = x / 10
	}
	fmt.Println(cantik)
}