package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
//
// [url: 请求地址 response: 请求返回的内容]
func Get(url string) (response string) {
	// 初始化HTTP客户端
	client := http.Client{Timeout: 5 * time.Second}
	// 发起请求获取响应
	resp, error := client.Get(url)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	// 解析响应内容
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		// 获取BODY
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	// Buffer 转换成字符串
	response = result.String()
	return
}

// 发送GET请求
//
// [header: 响应头 body: 响应体]
func GetURL(url string) (header http.Header, body *bytes.Buffer) {
	// 初始化HTTP客户端
	client := http.Client{Timeout: 5 * time.Second}
	// 发起请求获取响应
	resp, error := client.Get(url)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	// 获取 头部
	header = resp.Header

	// 解析 体部
	var buffer [512]byte
	body = bytes.NewBuffer(nil)
	for {
		// 获取BODY
		n, err := resp.Body.Read(buffer[0:])
		body.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return
}

// 发送POST请求
//
// url: 请求地址，data: POST请求提交的数据, contentType: 请求体格式
// content: 请求放回的内容
func Post(url string, data interface{}, contentType string) (content string) {
	// 数据JSON化
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// 设置POST , HTTP 头部
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// 获取HTTP客户端
	client := &http.Client{Timeout: 5 * time.Second}
	// 发起请求, 获取响应数据
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	// 解析响应数据
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
