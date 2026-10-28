package main

import "fmt"

func main() {
	var lamp_status string
	fmt.Print("Masukkan status lampu (merah/kuning/hijau): ")
	fmt.Scan(&lamp_status)
	language_translation := map[string]string{
		"merah":  "red",
		"kuning": "yellow",
		"hijau":  "green",
	}

	translated_status := language_translation[lamp_status]

	if translated_status == "red" {
		fmt.Println("Berhenti")
	}else if translated_status == "yellow" {
		fmt.Println("hati-hati")
	}else if translated_status == "green" {
		fmt.Println("jalan")
	}else {
		fmt.Println("Status lampu tidak valid")
	}

}

	