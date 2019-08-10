// 打开文件, 全部读取到内存 # os -> File -> Read / Write
// ioutil-> ReadFile(filename string) ([]byte, error) # 读取指定文件的全部内容, []byte 字节切片
// string(data []byte) # 字节切片转换成子字符串
// strings -> Split(s, sep string) []string # 字符串分割成子串切片
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// ioutil-> ReadFile(filename string) ([]byte, error) // 读取指定文件的全部内容
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		// 字符串分割成子串切片, "\n" 安装行分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
