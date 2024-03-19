package main

import (
	"fmt"
)

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	z := substrai(50, 12)
	a := divide(10, 2)
	fmt.Println(x, y, z, a)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total
}

func substrai(i ...int) int {
	if len(i) == 0 {
		return 0 // Se não houver argumentos, retorna 0
	}

	total := i[0] // Inicializa total com o primeiro elemento de i
	for _, v := range i[1:] {
		total -= v
	}

	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}

	total := i[0]
	for _, v := range i[1:] {
		total = total / v
	}

	return total
}
