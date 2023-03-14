package main

import (
	"fmt"
	"net"
)

func Tarsus() {
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
