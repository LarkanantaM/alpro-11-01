package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Print("Genap, ")
	} else {
		fmt.Print("Ganjil, ")
	}

	if n > 0 {
		fmt.Println("Positif")
	} else if n < 0 {
		fmt.Println("Negatif")
	} else {
		fmt.Println("Nol")
	}
}

//PSEUDOCODE
//Program Genap_Ganjil_Tanda
//Kamus
//  n : integer

//Algoritma
//baca(n)

//{ cek genap / ganjil }
//jika n mod 2 = 0 maka
//tulis("Genap, ")
//jika tidak
//tulis("Ganjil, ")
//akhirjika

//{ cek tanda bilangan }
//jika n > 0 maka
//tulis("Positif")
//jika tidak
//jika n < 0 maka
//tulis("Negatif")
//jika tidak
//tulis("Nol")
//akhirjika
//akhirjika
//Akhir Program