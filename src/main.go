package main

import (
	"fmt"
	"tarsus/go/src/structs"
)

func main() {
	ts := structs.NewTestStruct([]any{"1", "2"}).(structs.TestStruct)
	fmt.Printf("ts: %v\n", ts)
}
