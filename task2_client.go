package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var fileName = "data_task2-restore.txt" // 定义存储文件名
var mtx sync.Mutex                      // 写入文件需要互斥

func main() {
	fmt.Println("开始下载...")
	ch := make(chan string)
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// 启动3个协程并发下载分块数据
	go downloadBlock(file, "7777", ch)
	go downloadBlock(file, "8888", ch)
	go downloadBlock(file, "9999", ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Printf("下载完成 [%s]\n", fileName)
}

func downloadBlock(file *os.File, server string, ch chan<- string) {
	conn, err := net.Dial("tcp", ":"+server)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	mtx.Lock()

	// 创建buffer
	buf := make([]byte, 512)
	fileLenBuff := make([]byte, 8)

	// 先收文件大小
	_, _ = conn.Read(fileLenBuff)
	fileSize := int64(binary.BigEndian.Uint64(fileLenBuff))
	// 每块大小
	blockSize := fileSize / 3
	// 记录收到的字节数
	byteCount := 0

	if server == "7777" { // 收到的是第一块数据，来自主机1
		_, _ = file.Seek(0, 0)
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
			byteCount += cnt
		}

	} else if server == "8888" { // 收到的是第二块数据，来自主机2
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
			byteCount += cnt
		}

	} else if server == "9999" { // 收到的是第三块数据，来自主机3
		_, _ = file.Seek(blockSize*2, 0)
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
			byteCount += cnt
		}
	}
	mtx.Unlock()

	ch <- fmt.Sprintf("请求主机%s，成功下载了%d字节", server, byteCount)
	defer conn.Close()
}
