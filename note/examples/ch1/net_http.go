// HTTP 访问, 获取url资源
// http   -> Get(url string) (resp *Response, err error)
// http   -> type Response struct # 响应数据
//        -> Body io.ReadCloser   # 关闭输入流, 响应流
// ioutil -> ReadAll(r io.Reader) ([]byte, error) # 读取输入流, 返回字节切片
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		//
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v=\n", err)
			os.Exit(1)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		// 关闭输入流
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v=\n", err)
			os.Exit(1)
		}

		// 输入流数据
		fmt.Printf("%s", bytes)
	}

}
