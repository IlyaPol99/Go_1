package main

import "fmt"
import "math"

func main() {
	var a float64

	fmt.Println("Введите площадь круга:")
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Printf("Длина круга: % .2f\n", getCircleLength(a))
		fmt.Printf("Диаметр круга: % .2f\n", getCircleDiameter(a))
	} else {
		fmt.Println("Вы ввели площадь круга равную или меньше 0")
	}
}
func getCircleLength(a float64) float64 {
	return math.Sqrt(4 * math.Pi * a)
}
func getCircleDiameter(a float64) float64 {
	return math.Sqrt(4 * a / math.Pi)
}
