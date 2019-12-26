package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	fmt.Println("开始下载...")
	go spinner(100 * time.Millisecond)
	timestamp := time.Now().Format("15-04-05")
	fileName := "data_task1" + "-" + timestamp
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

// 下载过程动画
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
