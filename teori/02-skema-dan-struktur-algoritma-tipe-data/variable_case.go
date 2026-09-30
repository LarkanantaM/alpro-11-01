package main

import "fmt"

func main() {
	var name string 
	var age int
	var address string
	
	name = "Lingga Arkananta mahardika"	
	fmt.Println("Nama :", name)

	age = 19
	fmt.Println("Umur :", age)

	address = "Jl.Kanan lurus belok kiri 123,"
	fmt.Println("Alamat :", address)

	var middleName = "Arkananta"
	fmt.Println("Nama tengah :", middleName)

	lastName := "Mahardika"
	fmt.Println("Nama belakang :", lastName)

	var (
		fullName = "Lingga Arkananta Mahardika"
		firstName = "Mahardika"
	)
	fmt.Println(fullName)
	fmt.Println(firstName)
}