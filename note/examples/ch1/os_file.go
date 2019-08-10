// 打开文件, 按行读取
// os -> Open(name string) (*File, error) # 返回打开的文件指针, ERROR
// os -> (file *File) Close() error       # 关闭文件资源
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		// 没有文件参数
		// 安行读取标准命令行输入
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os -> Open(name string) (*File, error) // 返回打开的文件指针, ERROR
			// (file *File) Close() error // 关闭文件资源

			// bufio -> NewScanner(r io.Reader) *Scanner
			// bufio -> (s *Scanner) Scan() bool // 按行读取
			// bufio -> (s *Scanner) Text() string        // 获取内容
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
