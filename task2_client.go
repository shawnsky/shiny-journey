package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)



func main() {
	fmt.Println("开始下载...")
	ch := make(chan string)

	go downloadBlock("7777", ch)
	go downloadBlock("8888", ch)
	go downloadBlock("9999", ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("下载完成")
}

func downloadBlock(server string, ch chan<- string) {
	conn, err := net.Dial("tcp", ":"+server)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	// 文件不存在则创建
	fileName := "data_task2-restore"+server+".txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// 创建buffer
	buf := make([]byte, 512)
	fileLenBuff := make([]byte, 8)

	// 先收文件大小
	_, _ = conn.Read(fileLenBuff)
	fileSize := int64(binary.BigEndian.Uint64(fileLenBuff))
	// 每块大小
	blockSize := fileSize / 3

	fmt.Printf("Block size %d\n", blockSize)

	if server == "7777" {  // 收到的是第一块数据
		_, _ = file.Seek(0,0)
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

	} else if server == "8888" {  // 收到的是第二块数据
		_, _ = file.Seek(blockSize, 0)
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

	} else if server == "9999" {  // 收到的是第三块数据
		_, _ = file.Seek(blockSize * 2, 0)
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
	}

	ch <- fmt.Sprintf("下载完成 [%s]", fileName)
	defer conn.Close()
}
