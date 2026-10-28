package main

import "fmt"

func main()	{
	var N, mobil uint8
	fmt.Scan(&N)

	mobil = N / 7

	if N % 7 != 0 {
		mobil++
	}
	fmt.Println(mobil)
}