package main

import "fmt"

func main() {
	fmt.Printf("Hello World\n")

	a, b := 1, 2

	operationDone := make(chan bool)
	go func() {
		b = a * b

		operationDone <- true
	}()

	result := <-operationDone

	if result {
		a = b * b
	}

	fmt.Printf("a = %d, b = %d\n", a, b)
}