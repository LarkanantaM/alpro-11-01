package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan Nilai Anda: ")
		fmt.Scan(&nilai)
		if nilai >= 60 {
			fmt.Println("Anda Lulus Ujian")
		} else {
			fmt.Println("Anda Tidak Lulus Ujian")
		}

		if nilai == -1 {
			break
		}

	}
		fmt.Println("Program Selesai.")
}