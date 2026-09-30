package main

import "fmt"

func main() {
	const pi = 3.14
	
	//pi = 3.14159 // error: cannot assign to pi (constant value)

	fmt.Println("Nilai pi :", pi)
}