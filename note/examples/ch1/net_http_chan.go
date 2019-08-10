// HTTP 并发访问资源
// time -> Now() Time # 获取当前时间
// time -> Since(t Time) Duration # 计算时间t到当前时间时长
// time -> (d Duration) Seconds() float64 # 获取秒数

// ch := make(chan string) # 定义无缓冲通道, 值类型 string
// <-ch # 读通道, 阻塞
// ch <- string # 写通道, 阻塞

// go func -> 开启一个协程
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// 开始时间
	start := time.Now()

	// 定义通道
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// 开启协程获取URL资源
		go fetch(url, ch)
	}

	// 循环读取 ch, 堵塞进程
	for range os.Args[1:] {
		// 读通道
		fmt.Println(<-ch)
	}

	// 计算时长
	fmt.Printf("总耗时 -> %.2fs \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now() // 统计解析时间

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 获取资源错误, 退出
		return
	}

	// 拷贝输入流数据到 ioutil.Discard 丢弃
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		// 获取响应数据错误, 报告错误情况
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// 解析成功, 写入通道, 报告成功
	ch <- fmt.Sprintf("[耗时 %2.fs -> %7d 字节] [%s]", secs, nbytes, url)
}
