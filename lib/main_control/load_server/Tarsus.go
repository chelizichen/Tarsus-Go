package load_server

import (
	"fmt"
	"net"
	"strings"
)

/**
* 规定的协议枚举
 */
const (
	bufEndl   = "[#ENDL#]" // 结束标志
	protoEndl = "[##]"     // 协议结束标志
)

var (
	protoIndex = [5]string{"[#1]", "[#2]", "[#3]", "[#4]", "[##]"}
)

type protoPkg struct {
	id        uint
	interFace string
	method    string
	timeout   string
	data_len  string
	request   string
	data      []any
}

func LoadServer() {
	// 创建监听套接字
	listenSocket, err := net.Listen("tcp", ":3786")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	fmt.Println("TCP server is listening at :3786")

	// 无限循环，等待客户端连接
	for {
		// 等待客户端连接，并创建新的连接套接字
		conn, err := listenSocket.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr().String())

		// 新开一个协程处理客户端请求
		go handleClientRequest(conn)
	}
}

func LoadConfig() {

}

func handleClientRequest(conn net.Conn) {
	// 缓冲区大小为1024字节
	buf := make([]byte, 1024)

	for {
		// 读取客户端发送过来的数据，并放到缓冲区里面
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		// 将客户端发送过来的数据打印进行处理
		s := string(buf[0:n])
		pkg := unPackage(s)
		// 处理
		fmt.Println("Received message from client:", pkg.data)

		// 给客户端发送一个响应
		response := "Hello client!"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}

func unPackage(buf string) protoPkg {
	var pkg protoPkg
	proto := split(buf)
	pkg.interFace = getProto(proto[0], 1)
	pkg.method = getProto(proto[0], 2)
	//pkg.data = getData(proto[1])

	return pkg
}
func split(buf string) [2]string {
	index := strings.Index(buf, protoEndl)
	s1 := buf[0:index]
	s2 := buf[index+len(protoEndl) : len(buf)-len(bufEndl)-2]
	message := [2]string{s1, s2}
	return message
}

func getProto(proto string, index uint) string {
	// 截取上一个 加上 这一个 中间的值
	s1 := strings.Index(proto, protoIndex[index-1])
	s2 := strings.Index(proto, protoIndex[index])
	currProto := proto[s1+len(protoIndex[index-1]) : s2]
	return currProto
}
