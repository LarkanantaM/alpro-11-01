package main

import "fmt"

func main() {
	var angka, puluhan, satuan, hasil int
	fmt.Scan(&angka)
	puluhan = angka / 10
	satuan = angka % 10

	hasil = puluhan*1000 + puluhan*100 + satuan*10 + satuan
	fmt.Println(hasil)
}

// PSEUDOCODE
// Data yang sudah diketahui:
// angka adalah bilangan bulat positif yang terdiri dari dua digit
// Setiap digit akan digandakan
// Data yang dibaca:
// angka
// Proses:
// Mencari nilai digit puluhan dan satuan
// Menggandakan masing-masing digit
// Informasi yang dicetak:
// Hasil penggandaan digit

// ALGORITMA:
// {membaca nilai angka}
// INPUT angka
// {mencari digit puluhan}
// puluhan := angka DIV 10
// {mencari digit satuan}
// satuan := angka MOD 10
// {menggandakan masing-masing digit}
// hasil := puluhan*1000 + puluhan*100 + satuan*10 + satuan
// OUTPUT hasil