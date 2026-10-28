package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	total := 0
	first := true

	for x > 0 {
		digit := x % 10
		kuadrat := digit * digit
		total = total + kuadrat

		if first {
			fmt.Print(kuadrat)
			first = false
		} else {
			fmt.Print(" ", kuadrat)
		}
		x = x / 10
	}
	fmt.Println()
	fmt.Println(total)
}