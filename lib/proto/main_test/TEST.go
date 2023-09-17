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

func NewUser(intermediate []interface{}) interface{} {
	stream := proto.NewTarsusStream(intermediate, "User")
	fmt.Println(" 1 - ", stream.ReadString(1))
	fmt.Println(" 2 - ", stream.ReadString(2))
	fmt.Println(" 3 - ", stream.ReadString(3))
	fmt.Println(" 4 - ", stream.ReadString(4))
	fmt.Println(" 5 - ", stream.ReadString(5))
	return new(interface{})
}

func NewData(intermediate []interface{}) interface{} {
	stream := proto.NewTarsusStream(intermediate, "Data")
	fmt.Println(" 1 - ", stream.ReadString(1))
	fmt.Println(" 2 - ", stream.ReadString(2))
	fmt.Println(" 3 - ", stream.ReadStruct(3, "User"))
	fmt.Println(" 4 - ", stream.ReadList(4, "List<User>"))
	return new(interface{})
}

func main() {
	stream_data := "[\"0\",\"ok\",[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"],[\"1\",\"name\",\"11\",\"fullName\",\"address\"]]]"

	var intermediate []interface{}
	err := json.Unmarshal([]byte(stream_data), &intermediate)
	if err != nil {
		panic(err)
	}
	fmt.Println(stream_data)
	// Example usage

	proto.SetClass("User", proto.ClassBasic{
		Clazz:           reflect.TypeOf(User{}),
		Deserialization: NewUser,
	})
	proto.SetClass("Data", proto.ClassBasic{
		Clazz:           reflect.TypeOf(Data{}),
		Deserialization: NewData,
	})

	NewData(intermediate)
}
