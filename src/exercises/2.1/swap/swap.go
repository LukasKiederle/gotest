package swap

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Printf("Initial A: %d B: %d \n", a, b)

	swap2(&a, &b)
	fmt.Printf("Swapped A: %d B: %d \n", a, b)

	a, b = swap(a, b)
	fmt.Printf("Swapped Again A: %d B: %d \n", a, b)

	//a *int := nil

	//fmt.Printf("Nilpointer: %d", &a)

}

func swap(x, y int) (a, b int) {
	a, b = y, x
	return
}

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}
