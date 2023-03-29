package structs

import (
	"fmt"
	"tarsus/go/src/pkg"
)

type TestStruct struct {
	name string
	age  int
}

func NewTestStruct(args []any) any {
	var testStruct TestStruct
	pkg.StreamMap["TestStruct"] = NewTestStruct
	tarsusStream := pkg.TarsusStream(args, "TestStruct")
	fmt.Println(tarsusStream)
	testStruct.name = tarsusStream.ReadString(2)
	testStruct.age = tarsusStream.ReadInt(1)

	fmt.Println(testStruct.age)
	fmt.Println(testStruct.name)
	return testStruct
}
