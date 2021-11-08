package main

import (
	"fmt"
	"gb/lvl2/Lesson_2/Task2/isprime"
	"os"
)

func main() {
	var N int

	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scanln(&N); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for i := 2; i <= N; i++ {
		if isprime.IsPrime(i) {
			fmt.Print(i, " ")
		}
	}

	fmt.Print("\n")
}
