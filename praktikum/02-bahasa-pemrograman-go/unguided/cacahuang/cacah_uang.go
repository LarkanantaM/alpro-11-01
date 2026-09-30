package main
import "fmt"

func main() {
    var nominal int
    fmt.Scan(&nominal)

    lembar10ribu := nominal / 10000
    sisa := nominal % 10000

    lembar5ribu := sisa / 5000
    sisa = sisa % 5000
    
    lembar1ribu := sisa / 1000
  

    fmt.Print(lembar10ribu, " ", lembar5ribu, " ", lembar1ribu)
}