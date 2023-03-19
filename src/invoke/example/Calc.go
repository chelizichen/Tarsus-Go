package main

import (
	"fmt"
	"tarsus/go/src/invoke"
)

type Calculator struct {
	Num1 int
	Num2 int
}

func (c Calculator) Add(args int) int {
	return c.Num1 + c.Num2 +args
}

func main() {
	calculator := Calculator{Num1: 1, Num2: 2}
	result, err := invoke.Invoke(calculator, "Add",1)
	result1, err1 := invoke.Invoke(calculator, "Add",1)
	if err != nil || err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result[0].(int)) // 输出 3
	fmt.Println(result1[0].(int)) // 输出 3
}
