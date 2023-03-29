package main

import (
	"fmt"
	"tarsus/go/src/pkg"
)

type TestStruct struct {
	name string
	age  string
}

func NewTestStruct(args []any) any {
	var testStruct TestStruct
	pkg.StreamMap["TestStruct"] = NewTestStruct
	tarsusStream := pkg.TarsusStream(args, "TestStruct")
	fmt.Println(tarsusStream)
	testStruct.age = tarsusStream.ReadString(0)
	testStruct.name = tarsusStream.ReadString(1)
	fmt.Println(testStruct.age)
	fmt.Println(testStruct.name)
	return testStruct
}

func main() {
	NewTestStruct([]any{"1", "2"})
}
