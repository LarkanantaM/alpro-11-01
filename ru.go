package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	i := 1
	// REPEAT
	for {
		fmt.Printf("%d ", i)
		i++
		// UNTIL i > N
		if i > N {
			break
		}
	}
}