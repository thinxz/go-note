// HTTP WEB 服务
// http -> ListenAndServe(addr string, handler Handler) error # 创建 HTTP Server, 开始端口监听
// 每次处理请求, 都会启动新的 goroutine

// sync -> Mutex 互斥锁
// sync -> (m *Mutex) Lock()    # 获取锁
// sync -> (m *Mutex) Unlock()  # 释放锁

// fmt -> Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) # 格式化字符串, 输出到写入流

// io.Writer # 输出流接口
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 互斥锁
var mu sync.Mutex

var counts int

func main() {
	// 注册默认处理 Handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCount)
	http.HandleFunc("/info", handlerHead)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// 统计访问次数
func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // 获取锁
	counts++
	fmt.Fprintf(w, "Counts = %d\n", counts)
	mu.Unlock() // 释放锁
}

// HTTP Request -> head + form
func handlerHead(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		// 变量请求头
		_, _ = fmt.Fprintf(w, "header[%q], %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAdd = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
