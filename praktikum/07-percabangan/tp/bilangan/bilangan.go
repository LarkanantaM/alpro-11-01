package main 

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)


	if n == 0 {
		fmt.Print("Nol")	
	} else if n%2 == 0 {
		fmt.Print("genap, ")
		if n >= 1 {
			fmt.Print("positif")
		} else if n < 0 {
			fmt.Print("negatif")
		}
	} else if n%2 != 0 {
		fmt.Print("ganjil, ")
		if n >= 1 {
			fmt.Print("positif")
		} else if n < 0 {
			fmt.Print("negatif")
		}
	}

}