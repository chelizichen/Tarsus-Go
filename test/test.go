package main

import "strings"

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

const (
	buf_endl   = "[#ENDL#]" // 结束标志
	proto_endl = "[##]"     // 协议结束标志
)

func main() {
	i := []string{"111", "222"}
	println(i[0])
	println(i[1])
	//buf := "[#1]HelloInterFace[#2]TestRet[#3]3000[#4]36[##]#a##a#tom#b#jump#c#12#d#1#z##b##a#测试#z##c#1#z#[#ENDL#]"
	//of := indexOf(buf, "#", 5)
	//println(of)
	//message := split(buf)
	//getProto(message[0],1)
}

func split(buf string) [2]string {
	index := strings.Index(buf, proto_endl)
	s1 := buf[0:index]
	s2 := buf[index+len(proto_endl) : len(buf)-len(buf_endl)]
	message := [2]string{s1, s2}
	return message
}

func getProto(proto string, index uint) string {
	s1 := strings.Index(proto, proto_index[index-1])
	s2 := strings.Index(proto, proto_index[index])
	curr_proto := proto[s1+len(proto_index[index-1]) : s2]
	return curr_proto
}

// 测试通过
func indexOf(data string, find string, index uint) int {
	if index == 0 {
		curr_index := strings.Index(data, find)
		return curr_index
	}
	next_split := data[index:len(data)]
	curr_index := strings.Index(next_split, find)
	return curr_index
}
