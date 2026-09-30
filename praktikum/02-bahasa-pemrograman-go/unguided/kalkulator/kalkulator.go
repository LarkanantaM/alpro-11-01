package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("%d %d %d %d %d", a+b, a-b, a*b, a/b, a%b)
}