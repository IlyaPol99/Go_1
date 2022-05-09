package main

import (
	"fmt"
)

func main() {
	var a uint64
	fmt.Println("Введите положительное число")
	fmt.Scanln(&a)
	fmt.Printf("Результат нахождения числа Фибоначчи без буфера: %d\n", getFibonacci(a))
	fmt.Printf("Результат нахождения числа Фибоначчи с буфером: %d\n", getFibonacci2(a, make(map[uint64]uint64)))
}

// Фибоначчи без буфера
func getFibonacci(a uint64) uint64 {

	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return getFibonacci(a-1) + getFibonacci(a-2)
	}
}

// Фибоначчи с буфером
func getFibonacci2(a uint64, arr map[uint64]uint64) uint64 {
	if a == 0 {
		return a
	}
	if a == 1 {
		return a
	}
	prevNum := a - 1
	prevPrevNum := a - 2
	if _, found := arr[prevNum]; !found {
		arr[prevNum] = getFibonacci2(prevNum, arr)
	}
	if _, found := arr[prevPrevNum]; !found {
		arr[prevPrevNum] = getFibonacci2(prevPrevNum, arr)
	}
	return arr[prevNum] + arr[prevPrevNum]
}
