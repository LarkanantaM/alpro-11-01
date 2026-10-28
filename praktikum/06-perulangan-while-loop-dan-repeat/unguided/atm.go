package main

import (
	"fmt"

)

func main() {
	var pilihan, tarikSaldo int 
	var password int = 0
	var saldo int = 1000000
	n := 3

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&password)
		if password == 1234 {
			for {
				fmt.Println("Menu Atm: ")
				fmt.Println("1. Lihat Saldo")
				fmt.Println("2. Tarik Uang")
				fmt.Println("3. Keluar")			
				fmt.Print("Pilih opsi: ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {
					fmt.Println("Saldo Anda: ",saldo)
				} else if pilihan == 2 {
					fmt.Print("Masukkan nilai yang ingin anda tarik: ")
					fmt.Scan(&tarikSaldo)
					fmt.Println("Anda berhasil menarik sebanyak: ",tarikSaldo)
					saldo = saldo - tarikSaldo
					fmt.Println(saldo)
				} else if pilihan == 3 {
					fmt.Println("Keluar dari sistem.")
					break
				} else {
					fmt.Println("Pilihan tidak valid")
					break
				}
			}
			break

		} else {
			if i == n {
				fmt.Print("Rekening anda telah blokir.")
			} else {
				fmt.Printf("Password salah. Silahkan coba lagi.\n")
			}
		}


	}
}