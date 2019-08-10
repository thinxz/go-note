// 命令行参数 [command-len arguments] -> 安行读取标准命令行输入
// os 包, 提供跨平台 与操作系统交互的函数和变量
// os.Args    # 字符串切片值
// os.Args[0] # 命令本身名称
// os.Args[1;len(os.Args)] # 传入参数切片
// for -> 遍历区间 range #  range => 索引, 对应索引元素的值
package main

import (
	"fmt"
	"os"
)

// os.Args 字符串切片值
// os.Args[0] 命令本身名称
// os.Args[1;len(os.Args)] 传入参数切片
func main() {
	var s, sep string
	//
	for i := 1; i < len(os.Args); i++ {
		// 连接字符 : sep + os.Args[i]
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	for _, arg := range os.Args[1:] {
		// 遍历区间 range #  range => 索引, 对应索引元素的值
		fmt.Println(arg)
	}
}
