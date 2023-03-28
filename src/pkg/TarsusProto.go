package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

type any = interface{}

type proto_pkg struct {
	id        uint
	interFace string
	method    string
	timeout   string
	data_len  string
	request   string
	data      []any
}

var (
	proto_index = [5]string{"[#1]", "[#2]", "[#3]", "[#4]", "[##]"}
)

/**
* 规定的协议枚举
 */
const (
	buf_endl   = "[#ENDL#]" // 结束标志
	proto_endl = "[##]"     // 协议结束标志
)

/*
*

	@ReturnType [...]string
*/
func unPackage(buf string) proto_pkg {
	var pkg proto_pkg
	proto := split(buf)
	pkg.interFace = getProto(proto[0], 1)
	pkg.method = getProto(proto[0], 2)
	//pkg.data = getData(proto[1])

	return pkg
}
func split(buf string) [2]string {
	index := strings.Index(buf, proto_endl)
	s1 := buf[0:index]
	s2 := buf[index+len(proto_endl) : len(buf)-len(buf_endl)-2]
	message := [2]string{s1, s2}
	return message
}

func getProto(proto string, index uint) string {
	// 截取上一个 加上 这一个 中间的值
	s1 := strings.Index(proto, proto_index[index-1])
	s2 := strings.Index(proto, proto_index[index])
	currProto := proto[s1+len(proto_index[index-1]) : s2]
	return currProto
}

type User struct {
	Name string
	Age int
}

// JSON DEMO
func getData(buf string) User{
	jsonStr := `{"name":"John","age":30}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name, user.Age)
	return user;
}
