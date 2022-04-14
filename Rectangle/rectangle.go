package main

import "fmt"

func main() {
	var a uint
	var b uint

	fmt.Println("Введите длину первой стороны прямоугольника, целое положительное число:")
	fmt.Scanln(&a)
	fmt.Println("Введите длину второй стороны прямоугольника, целое положительное число:")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника: %d\n", a*b)
}
