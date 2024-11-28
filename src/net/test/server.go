package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [256]byte
		// 读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("reader.Read  error : ", err)
			break
		}
		recvData := string(buf[:n])
		fmt.Println("receive data ：", recvData)
		// 将数据再发给客户端
		conn.Write([]byte(recvData))
	}
}

func main() {
	// 监听tcp
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen  error : ", err)
		return
	}
	for {
		// 建立连接 ， 看到这里的朋友，有没有觉得这里和C/C++的做法一毛一样
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error : ", err)
			continue
		}
		// 专门开一个goroutine去处理连接
		go process(conn)
	}
}
