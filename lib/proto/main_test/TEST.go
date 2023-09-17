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
	inst := new(User)
	stream := proto.NewTarsusStream(intermediate, "User")
	inst.Id = stream.ReadString(1)
	inst.Name = stream.ReadString(2)
	inst.Age = stream.ReadString(3)
	inst.FullName = stream.ReadString(4)
	inst.Address = stream.ReadString(5)
	return inst
}

func NewData(intermediate []interface{}) interface{} {
	inst := new(Data)
	stream := proto.NewTarsusStream(intermediate, "Data")
	inst.Code = stream.ReadString(1)
	inst.Message = stream.ReadString(2)
	inst.Data = *stream.ReadStruct(3, "User").(*User)
	inst.Users = stream.ReadList(4, "List<User>").([]User)
	return inst
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
		Clazz:           reflect.TypeOf(&User{}),
		Deserialization: NewUser,
	})
	proto.SetClass("Data", proto.ClassBasic{
		Clazz:           reflect.TypeOf(&Data{}),
		Deserialization: NewData,
	})

	data := NewData(intermediate).(*Data)
	fmt.Println(data.Users)
}
