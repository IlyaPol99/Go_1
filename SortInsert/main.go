package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := make([]int64, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите числа для создаваемого массива. По окончании введите не число.\n")
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("Начальный массив:\n", arr)
			break
		} else {
			arr = append(arr, num)
		}
	}
	sortByInsert(arr)
	fmt.Println("Отсортированный массив:\n", arr)
}

func sortByInsert(arr []int64) {
	for i := 1; i < len(arr); i++ {
		nextEl := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > nextEl; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = nextEl
	}
}
