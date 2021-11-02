package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type MyError struct {
	errorTime time.Time
	errorText string
}

func (e *MyError) Error() error {
	return errors.New(e.errorTime.Format("2006-01-02 15:04:05") + ": " + e.errorText)
}

func NewMyError() *MyError {
	return &MyError{}
}

var err = NewMyError()

func main() {
	var result int

	result = divide(1, 0)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}

func divide(a, b int) int {
	defer func() {
		if e := recover(); e != nil {
			err.errorTime = time.Now()
			err.errorText = "А-a-a-a!!! Dividing by ZERO!!! I'm in a panic!"
		}
	}()

	return a / b
}
