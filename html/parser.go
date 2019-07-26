// HTML DOM 树解析
package html

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// 定义结点[子结点, 可以为Element或string]
type Node interface{}

// 定义结点元素
type Element struct {
	// 标签名
	tag string
	// 标签属性 [名称 + 值]
	attrs []xml.Attr
	// 子结点 [元素或字符值]
	children []Node
}

// HTML 解析 [基于XML解析器]
//
// 返回DOM, 结点树
// 结点名 [属性] [子元素]
// 属性 => {{属性名} 值}
//& {html [{{lang} en}] [
//	    {head []
//	        {meta [{{ charset} UTF-8}] []}
//	        {meta [{{ name} viewport} {{ content} width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0}] []}
//	        {meta [{{ http-equiv} X-UA-Compatible} {{ content} ie=edge}] []}
//	        {title [] [Document]}
//	     ]}
//       {body [] [
//           {h1 [{{ name} haodawang}] [Hello World]}
//        ]}
//  ]}
func Parser(reader io.Reader) (*Element, error) {
	// XML 解析器
	decoder := xml.NewDecoder(reader)

	// 01 声明一个 Element指针类型的栈
	// 第一是防止栈的空间随数据的膨胀成正相关，第二是后面涉及到修改元素的 children
	var stack []*Element
	// 设置 Element 的 children 让所有 Element 都关联成一颗树
	// 声明一个 currentNode 的 *Element 类型变量，来保存当前的 Element

	var currentElement *Element
	for {
		//
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}

		// discriminated union 也就是可辨识联合，通过 switch 去断言当前的 token 是哪个类型，然后 dispatch 相应的处理
		switch token := token.(type) {
		case xml.StartElement:
			// 遇到开始标签, push 进栈
			// StartElement类型的时候就 push 进栈
			stack = append(stack,
				&Element{tag: token.Name.Local, attrs: token.Attr, children: []Node{}})
			break
		case xml.EndElement:
			// 遇到闭合标签，pop 出栈
			// 遇到 EndElement 的时候就把它 pop 出来，这样就能组成一个完成的标签元素
			currentNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			preNode := stack[len(stack)-1]
			preNode.children = append(preNode.children, *currentNode)
			currentElement = preNode
			break
		case xml.CharData:
			// 标签中的字符串内容，解析到之后直接 push 进栈中顶层元素的 children 中
			if len(stack) == 0 {
				break
			}
			lastNode := stack[len(stack)-1]
			lastNode.children = append(lastNode.children, string(token[:]))
			break
		}
	}

	return currentElement, nil
}
