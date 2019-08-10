# go lang syntax

``` 表达式 [保留字 / 运算符 / 初始化 / 控制流]


```

``` 01 保留字

  if        else        switch      select      case
  
  for       range       default     break       continue

  return    var         const       type        struct

  func      interface   defer       map         chan

  import    package     go          goto        fallthrough

```

``` 02 运算符

  # i++ // 自增语句
  # i-- // 自减语句


```

``` 03 初始化


```

``` 04 控制流

  // GO 只有 for 循环一种
  for initialization; condition; post {
    // initialization 循环开始前执行, 简单语句 [短变量声明, 自增语句, 赋值语句, 函数调用]
    // condition 布尔表达式, 每次循环迭代开始前判断 true 执行
    // post      循环体执行结束后执行
    // zero or more statements
  }
  
  // for 遍历区间 range #  range => 索引, 对应索引元素的值
  for _, arg := range os.Args[1:] {
  // _ 空标识符
    fmt.Println(arg)
  }
  
  // 条件判断语句
  if condition {
  
  }
  
  //
  if condition {
  } else {
  }

```