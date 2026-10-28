package main

import "fmt" 

func main() {
	var a int
	fmt.Scan(&a)
	jumlah :=  0

	for i := 1; i <= a; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			jumlah = jumlah + i
		}
	}
	fmt.Println(jumlah)
}