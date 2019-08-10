// os.Stdin # 标准输入流 -> io.Reader
// bufio -> NewScanner(r io.Reader) *Scanner # 创建对象 Scanner 类型, 读取数据并将其拆成行或单词
// bufio -> (s *Scanner) Scan() bool      # 按行读取, true， 无读取输入时返回 false
// bufio -> (s *Scanner) Text() string    # 获取读取内容

// map 字典结构 -> map[key] value  # value := map[key] 获取值
// make(make[string]int) 创建map结构变量并分配内存, key : string , value : int
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 内置函数 make()
	// 字典类型 map
	counts := make(map[string]int)

	delete(counts)

	// bufio.NewScanner(os.Stdin) , 标准输入流 创建对象 Scanner 类型 读取数据并将其拆成行或单词
	// input.Scan() 读到一行返回 true， 无输入时返回 false
	// input.Text() 获取内容
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 循环输入
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			// 格式化输出
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
