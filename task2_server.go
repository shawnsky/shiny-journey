package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var filename = "data_task2.txt"

func sendBlockedFileHandler(c net.Conn) {
	defer c.Close()
	// 打开目标文件
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	// 获取文件长度
	fileInfo, _ := os.Stat(filename)
	fileSize := fileInfo.Size()
	// 计算需要读满几次buff
	buffCount := fileSize / 512
	// 计算每块数据需要读几次buff
	sendCount := int(buffCount / 3)
	// 创建buffer
	buf := make([]byte, 512)

	flag := os.Args[1]

	if flag == "7777" { // 发送第一块数据
		for i := 0; i < sendCount; i++ {
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
	} else if flag == "8888" { // 发送第二块数据
		sent := int64(sendCount) * 512
		// 移动文件指针
		_, _ = fs.Seek(sent, 0)
		for i := 0; i < sendCount; i++ {
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

	} else if flag == "9999" {
		sent := int64(sendCount) * 512 * 2
		// 移动文件指针
		_, _ = fs.Seek(sent, 0)
		for i := 0; i < sendCount; i++ {
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
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请使用端口号7777/8888/9999作为参数")
	}
	port := os.Args[1]

	if port != "7777" && port != "8888" && port != "9999" {
		fmt.Println("启动失败，端口号有误")
		return
	}
	// 建立对应端口号服务器和主机id的关联
	// 文件一共分为3块，主机id为1，代表提供第一块数据
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("fail to listen")
		os.Exit(1)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("fail to accept")
		}
		go sendBlockedFileHandler(conn)
	}
}
