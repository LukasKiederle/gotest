package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100  //huge
	Small = Big >> 99 //2
)

var c, python, java bool

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Printf("This is pi: %g \n", math.Pi)

	fmt.Printf("Result of adding %d and %d is: %d \n", 10, 15, add(10, 15))

	fmt.Printf("Normal: %s %s \n", "test1", "test2")
	a, b := swap("test1", "test2")
	fmt.Printf("Inverted: %s %s \n", a, b)

	x, y := split(10)
	fmt.Printf("Result of splitting 10 is: %d and %d \n", x, y)

	fmt.Println(c, python, java)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func split(sum int) (x, y int) {
	x = sum / 4
	y = sum - x
	return
}

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
