package main

import "fmt"

// どちらも同じ型であれば、型を省略できる
func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	}

	return a / b, nil
}
