# go lang syntax

``` 类型 [变量 / 常量 / 基本类型 / 引用类型 / 类型转换 / 字符串 / 指针 / 自定义类型]

  GO 编译型语言; GO 原生支持 Unicode;
  
  # go run xxx.go    // 直接运行
  # go build xxx.go  // 编译
  # go get xxx.com/xx/xx // 自动获取仓库源码

  # 包概念 [类型, 库 libraries 或者模块 modules] -> package import
  
  # gofmt
  # go get golang.org/x/tools/cmd/goimports

```

``` 01 变量

  GO 静态类型语言, 不能在运行期改变变量类型
  
  A、var varName int # var 变量名 基本类型
  var varName1, varName2 int
  
  B、x := 10 # 变量名 := 值 # 短声明变量, 只适用于局部变量
  x, y := 10, 10

  C、
  var(
    x int
    y int
  )

```

``` 02 常量

  A、const x, y int = 1, 2 # const 常量名 常量类型 = 值

  B、
  const(
    a, b   = 10, 100
    c bool = false
  )
  
  C、
  const(
    s = "AAA"
    x      // 未提供类行和初始化, 类同上一个常量 => x = "AAA"
  )
  
  D、枚举 iota 定义常量足, 从0开始计数, 自增枚举值
  const(
    Sundy = iota // 0
    Monday       // 1
    Tuesday
    Wendnesday
    Thursday
    Friday
    Saturday
  )

```

``` 03 基本类型 [Unicode]


  byte          1B   0       # uint8
  rune          4B   0       # unicode code point, int32
  int,uint      4/8B 0       # 32/64位
  int8,uint8    1B   0       # -128 - 127, 0 - 255
  int16, uint16 2B   0       # -32768 - 32767, 0 - 65535
  int32, uint32 4B   0       # -21亿 - 21亿, 0 - 42亿
  int64, uint64 8B   0
  float32       4B   0.0
  float64       8B   0.0
  complex64     8B
  complex128    16B
  uintptr       4/8B         # 存储指针的 uint32 / uint64

  array                      # 值类型
  struct                     # 值类型

  bool          1B   false
  string             ""      # UTF-8

  slice              nil     # 引用类型
  map                nil     # 引用类型
  channel            nil     # 引用类型
  interface          nil     # 接口
  function           nil     # 函数

  进制 => [八进制 071 / 十六进制 0x1F / 科学计数 1e9]
```

``` 04 引用类型 [slice / map / channel , 包含复杂的内部结构, 需要申请内存 以及初始化相关属性]

  => 内置函数 new , 计算类型大小并分配零值内存, 返回指针
  => 内置函数 make, 被编译器翻译成具体的创建函数, 由其分配内存和初始化成员结构, 返回对象而非指针
 
  A、 x := []int{0, 0, 0}  // 提供初始化表达式
  x[0] = 0

  B、 x := make([]int, 3)  // makeslice
  x[0] = 0
  
  C、 x := new([]int)
  x[0] = 0                 // Error

```

``` 05 类型转换 [不支持隐式类型转换]

  A、 var b byte = 1
  var x int = int(b)  # 显示转换

```

```  06 字符串 [字符串为不可变类型]

  # + , 字符连接符号, 连接两个字符串为新字符串 => s = "A" + "B"


```

```  07 指针


```

```  08 自定义类型

  类型 : 命名类型 [bool, int, string ...] 和 未命名类型 [array, slice, map ...]
  相同声明的未命名类型视为同一类; 重新定义的命令类型不相同, 必须显示转换

  => type 新类型名  类型名 # type bigint int64 定义新类型
  x := 0
  var y bigint = bigint(x) # 必须显示转换 (不包含常量)

  var s newslice = []int{1, 2, 3}
  var s1 []int = s                # 未命名类型, 隐式转换

```

·········· ·········· ·········· ·········· ··········

``` 数据 [Array / Slice / Map / Struct]


```

``` 01 Array

  # 值类型, 复制传参会复制整个数组
  # 数组长度为常量, 且是类型的组成部分
  
  => 内置函数 len, 返回数组长度
  => 内置函数 cap, 返回数组长度

```

``` 02 Slice 切片

  # 引用类型, 自身是结构体, 值拷贝传递
  # 属性 len , 可用元素数量
  # 属性 cap , 最大扩张容量
  
  => reslice , 基于原slice 创建新 slice对象
  => append  , slice 尾部添加数据, 返回新 slice 对象
  => copy
  
  # s[i] 访问第i的单个元素
  # s[m:n] 获取子序列    // 产生从第m个元素到第n-1个元素的切片
  # 省略m 或 n, 默认传入 0 或 len(s) -> s[:n] => s[0:n] 、s[0:] => s[0:len(s)]
  # len(s) 获取元素数目

```

``` 03 Map 字典

  # 引用类型，哈希表 - key/value 键任意类型, 但需要能 == 运算符比较; 值任意类型
  # map[key] value => m := make(map[string]int)
  # m[key] = value / m[key] 不存在时, 自动初始化对应类型的零值

```

``` 04 Struct

  # 值类型


```

``` 面向对象

  # GO 仅支持封装
  # 匿名字段的内存布局和行为类型继承

```















