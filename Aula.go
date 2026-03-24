package main

import "fmt"

func main() {
	var n1, p1, n2, p2 float64
	fmt.Println("Digite a nota:")
	fmt.Scan(&n1)
	fmt.Println("Digite o peso da nota:")
	fmt.Scan(&p1)
	fmt.Println("Digite sua segunda nota:")
	fmt.Scan(&n2)
	fmt.Println("Digite o peso da nota:")
	fmt.Scan(&p2)
	nota1 := n1 * p1
	nota2 := n2 * p2
	somar := nota1 + nota2
	peso := p1 + p2
	final := somar / peso
	fmt.Println("Sua media final é:", final)
}
