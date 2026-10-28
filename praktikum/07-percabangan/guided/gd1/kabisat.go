package main

import "fmt"

func main() {
	var tahun int
	fmt.Scanln(&tahun)

	if (tahun % 400 ==0) || (tahun % 4 == 0 && tahun % 100 != 0) {
		fmt.Println("Kabisat")
	} else {
		fmt.Println("Bukan Kabisat")
	}
}