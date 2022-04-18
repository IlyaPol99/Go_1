package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var resF uint64
	var f int
	var op string
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	if op == "!" {
		fmt.Print("Введите положительное число: ")
		fmt.Scanln(&f)
	} else {
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Print("Делить на ноль нельзя\n")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "!":
		resF = getFactorial(f)
	default:
		fmt.Println("Операция выбрана неверно\n")
		os.Exit(1)
	}
	if op == "!" {
		fmt.Printf("Результат выполнения операции: %d\n", resF)
	} else {
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	}
}

func getFactorial(f int) uint64 {
	var res uint64 = 1
	if f < 0 {
		fmt.Print("Отрицательный факториал не существует\n")
		os.Exit(1)
	} else {
		for i := 1; i <= f; i++ {
			res = res * uint64(i)
			fmt.Println(res)
		}
	}
	return res
}
