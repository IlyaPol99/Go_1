package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num int

	fmt.Fprintln(os.Stdout, "Введите целое положительное число:")
	fmt.Fscanln(os.Stdin, &num)
	if num < 2 {
		fmt.Fprintln(os.Stdout, "Простые числа не найдены")
		return
	} else {
		fmt.Fprintln(os.Stdout, "Простые числа:")
	}

	arr := make([] int,(num+1))
	limitation := int(math.Sqrt(float64(num)) + 1)

	// Убираем чётные
	for i := 4; i <= num; i++ {
		if i%2 == 0 {
			arr[i] = 1
		}
	}
	for i := 3; i < (limitation); i++ {
		if arr[i] == 0 {
			for j := i * i; j <= num; j = j + i {
				arr[j] = 1
			}
		}
	}
	// Вывод чисел
	for i, _ := range arr {
		if i > 1 && arr[i] == 0 {
			fmt.Printf("%d\n",i)
		}
	}
}
