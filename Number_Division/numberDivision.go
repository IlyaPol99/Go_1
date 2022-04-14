package main

import "fmt"

func main() {
	var a uint

	fmt.Println("Введите трёхзначное целое положительное число:")
	fmt.Scanln(&a)

	fmt.Printf("Единицы числа %d: %d\n", a, a%10)
	fmt.Printf("Десятки числа %d: %d\n", a, a/10%10)
	fmt.Printf("Сотни числа %d: %d\n", a, a/100%10)
}
