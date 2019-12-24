package main

import (
	"fmt"
	"io"
	"net"
	"os"
)


func main() {
	fmt.Println("开始下载...")
	downloadBlock("7777")
	downloadBlock("8888")
	downloadBlock("9999")
}

func downloadBlock(server string) {
	conn, err := net.Dial("tcp", ":"+server)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	fileName := "data_task2-" + server + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// 创建buffer
	buf := make([]byte, 512)
	// 循环读数据并写入文件
	for {
		cnt, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if cnt == 0 {
			break
		}
		if _, err := file.Write(buf[:cnt]); err != nil {
			panic(err)
		}
	}
	fmt.Printf("下载完成 [%s]\n", fileName)
	defer conn.Close()
}
