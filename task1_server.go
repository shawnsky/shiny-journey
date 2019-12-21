package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var filename = "data"

func connHandler(c net.Conn) {
	defer c.Close()
	// 打开目标文件
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	// 创建buffer
	buf := make([]byte, 512)
	// 循环发送文件数据
	for {
		cnt, err := fs.Read(buf)
		if err == io.EOF {
			break
		}
		_, err = c.Write(buf[:cnt])
		if err != nil {
			fmt.Println("文件发送失败，可能是客户端已断开")
			break
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println("fail to listen")
		os.Exit(1)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("fail to accept")
		}
		go connHandler(conn)
	}
}