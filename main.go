package main

import "fmt"

func main() {
	x := soma(1, 2, 3, 4, 5)
	y := subtrai(1, 2, 3, 4, 5)
	w := multiplica(1, 2, 3, 4, 5)
	z := divide(1, 2, 3, 4, 5)
	fmt.Println(x, y, w, z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total -= v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
