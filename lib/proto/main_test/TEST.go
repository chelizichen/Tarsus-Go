package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"tarsus/go/lib/proto"
)

type User struct {
	Id       string
	Name     string
	Age      string
	FullName string
	Address  string
}

type Data struct {
	Code    string
	Message string
	Data    User
	Users   []User
}

func main() {
	stream_data := "[\"0\",\"ok\",[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"]]]"

	var intermediate []interface{}
	err := json.Unmarshal([]byte(stream_data), &intermediate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", stream_data)
	fmt.Println()
	// Example usage
	proto.SetClass("User", reflect.TypeOf(User{}))
	proto.SetClass("Data", reflect.TypeOf(Data{}))

	stream := proto.NewTarsusStream(intermediate, "Data")
	fmt.Printf(" 1 - %s", stream.ReadString(1))
	fmt.Printf(" 2 - %s", stream.ReadString(2))
	fmt.Printf(" 3 - %s", stream.ReadStruct(3, "User"))
}
