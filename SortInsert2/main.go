package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int64, 0, 10)
	var inArr string
	fmt.Println("Введите числа через пробел для создаваемого массива.\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inArr = scanner.Text()
		res := strings.Split(inArr, " ")
		if len(res) > 0 {
			for i := 0; i < len(res); i++ {
				num, err := strconv.ParseInt(res[i], 10, 64)
				if err != nil {
					return
				} else {
					arr = append(arr, num)
				}
			}
			sortByInsert(arr)
			fmt.Println("Отсортированный массив:\n", arr)
		}
		return
	}
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
