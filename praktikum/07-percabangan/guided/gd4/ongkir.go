package main 

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int 
	var layanan string
	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebhihi batas maksimum!")
		return
	}

	var tarif int

	switch zona {
	case 1:
		tarif = 10000
	case 2:
		tarif = 15000
	case 3:
		tarif = 25000
	default :
		fmt.Print("Zona tidak valid!")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarif

	switch layanan {
	case "Reguler":
		fmt.Print()
	case "Express":
		total += total/2
	default:
		fmt.Println("Tidak valid!")
		return
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}