package main

import (
	"fmt"
	"math/rand"
	"time"
)

var N = 10         // 启用的协程数量
var each = 1000000 // 每个协程计算一百万个点
var r = 100.0
var countIn = 0
var countAll = each * N

func calculate(ch chan<- int) {
	beginTime := time.Now()
	var in = 0
	for i := 0; i < each; i++ {
		x := genRandomFloat64(-100, 100)
		y := genRandomFloat64(-100, 100)
		if x*x+y*y < r*r {
			in += 1
		}
	}
	cost := time.Since(beginTime).Milliseconds()
	fmt.Printf("任务耗时 %dms\n", cost)
	ch <- in
}

func main() {
	ch := make(chan int)

	rand.Seed(time.Now().UnixNano())

	beginTime := time.Now()

	for i := 0; i < N; i++ {
		go calculate(ch)
	}

	for i := 0; i < N; i++ {
		countIn += <-ch
	}

	pi := float64(countIn) / float64(countAll) * 4
	cost := time.Since(beginTime).Milliseconds()

	fmt.Printf("CountIn=%d CountAll=%d, PI=%f, 总耗时%dms", countIn, countAll, pi, cost)
}

func genRandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
