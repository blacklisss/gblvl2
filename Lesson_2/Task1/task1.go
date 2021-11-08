package main

import (
	"fmt"
)

func main() {

	divide(1, 0)
	fmt.Println("Don't Panic! We survived dividing by zero!")
}

func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A-A-a-a-a!!! Dividing by ZERO!!! I'm in a panic!!!")
		}
	}()

	return a / b
}
