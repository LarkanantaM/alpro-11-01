package main

import "fmt"

func main() {
	var total float64
	var asesmen bool

	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&total)
	fmt.Print("Sedang asesmen CLO? (true/false): ")
	fmt.Scan(&asesmen)

	var totalAkhir float64

	if asesmen == true {
		totalAkhir = total - (total * 35 / 100)
	} else {
		totalAkhir = total
	}

	fmt.Printf("Total belanja akhir yang harus dibayarkan: %.0f\n", totalAkhir)
}