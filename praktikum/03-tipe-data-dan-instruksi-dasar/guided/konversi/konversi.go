package main

import "fmt"

func main()	{
	var celcius float64;

	fmt.Println("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celcius)
	
	fmt.Println(celcius + 273)
}