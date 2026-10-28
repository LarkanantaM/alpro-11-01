package main

import "fmt"

func main() {
	var a, b int
	var op string
	fmt.Scanln(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Division by zero")
		} else {
			fmt.Printf("%.2f", float64(a)/float64(b))
		}
		case "%":
		if b == 0 {
			fmt.Println("Error: Division by zero")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Error: Invalid operator")
	}
}
