package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	hasil := 0

	for i := 1; i <= n; i++{
		hasil += i
	}
	fmt.Println(hasil)
}