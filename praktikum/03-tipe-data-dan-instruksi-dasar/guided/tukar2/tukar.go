package main

import "fmt"

func main() {
	var x, y, z, temp int
	fmt.Scan(&x, &y, &z)

	temp = x
	x = y
	y = z
	z = temp

	fmt.Println(x, y, z)
}