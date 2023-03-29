package structs

import (
	"tarsus/go/src/pkg"
)

type TestStruct struct {
	name string
	age  int
}

func NewTestStruct(args []any) any {
	var testStruct TestStruct
	tarsusStream := pkg.TarsusStream(args, "TestStruct")
	testStruct.name = tarsusStream.ReadString(1)
	testStruct.age = tarsusStream.ReadInt(2)
	return testStruct
}

func init() {
	pkg.StreamMap["TestStruct"] = NewTestStruct
}
