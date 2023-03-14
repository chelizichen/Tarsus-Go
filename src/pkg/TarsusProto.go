package pkg

import (
	"strings"
)

type any = interface{}

type proto_pkg struct {
	id        uint
	interFace string
	method    string
	timeout   string
	data_len  string
	data      []any
}

var (
	proto_index = [5]string{"[#1]", "[#2]", "[#3]", "[#4]", "[##]"}
	size        = [26]string{
		"#a#", "#b#", "#c#", "#d#",
		"#e#", "#f#", "#g#", "#h#",
		"#i#", "#j#", "#k#", "#l#",
		"#m#", "#n#", "#o#", "#p#",
		"#q#", "#r#", "#s#", "#t#",
		"#u#", "#v#", "#w#", "#x#",
		"#y#", "#z#",
	}
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

	pkg.data = getData(proto[1])

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
	s1 := strings.Index(proto, proto_index[index-1])
	s2 := strings.Index(proto, proto_index[index])
	curr_proto := proto[s1+len(proto_index[index-1]) : s2]
	return curr_proto
}

func getData(buf string) []any {
	ret := make([]any, 0)
	init := 0
	start := strings.Index(buf, size[init])

	// 相当于while
	for {
		next_init := init + 1
		next := indexOf(buf, size[next_init], start)
		if next == -1 {
			// 是否分割完
			if start+len(size[init]) == len(buf) {
				break
			}
			// 切片查看是否是没切完的
			sub_pkg := buf[start : start+6]
			is_un_pkg := sub_pkg == size[init]+size[0]
			if is_un_pkg {
				un_pkg := buf[start+3 : len(buf)-3]
				args := getData(un_pkg)
				ret = append(ret, args)
			} else {
				arg := buf[start+3 : len(buf)-3]
				ret = append(ret, arg)
			}
			break
		} else {
			isObject := buf[start:start+6] == size[init]+size[0]
			if isObject {
				curr_end_str := size[len(size)-1] + size[init+1]
				end := strings.Index(buf, curr_end_str)
				un_pkg := buf[start+3 : end+3]
				args := getData(un_pkg)
				ret = append(ret, args)
				start = end + 3
			} else {
				println(start, " --- ", next, " --- ", len(buf))
				arg := buf[start+3 : next]
				println(arg)
				ret = append(ret, arg)
				start = next
			}
		}
		init++
	}
	return ret
}
