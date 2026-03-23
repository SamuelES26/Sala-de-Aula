package main

import "fmt"

func main() {
	var ld1, ld2, ld3 int64
	fmt.Println("Coloque os dados:")
	fmt.Scan(&ld1, &ld2, &ld3)
	if ld1 == ld2 && ld2 == ld3 {
		fmt.Println("Equilátero")
	} else if ld1 == ld2 || ld1 == ld3 || ld3 == ld2 {
		fmt.Println("Isosceles")
	} else if ld1 != ld2 && ld2 != ld3 || ld2 != ld1 && ld2 != ld3 {
		fmt.Println("Escaleno")
	}
}
