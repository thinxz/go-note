package main

import (
	"flag"
	"fmt"
	"github.com/thinxz/go_lang/html"
	"os"
)

// HTML 文档所在目录
var path string

func init() {
	// 注册参数
	flag.StringVar(&path, "c", "", "")
}

func main() {

	// 解析参数
	flag.Parse()
	htmlFile := path + "/index.html"
	//
	fmt.Printf("HTML Path : %s", htmlFile)

	// 打开文件
	file, err := os.Open(htmlFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// 解析 HTML 文档
	ele, err := html.Parser(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// 输出解析结果
	fmt.Println(ele)
}
