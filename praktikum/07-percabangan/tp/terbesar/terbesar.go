package main 

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if x > y {
		fmt.Println(x)
	} 
	if y > x {
		fmt.Println(y)
	} 
	if z > y {
		fmt.Println(z)
	}
}