// REDIS 客户端, 订阅消息并解析处理
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/thinxz/go_lang/config"
	"io"
	"log"
	"strings"
)

var (
	redisClient *redis.Client // REDIS客户端
	queue       chan string   // 队列通道
)

// 消息对象定义
// ---------- ---------- ---------
// Type    消息类型
// No      消息标识符
// Origin  消息内容
// ---------- ---------- ---------
type Message struct {
	Type   int    `json:"type"`
	No     string `json:"no"`
	Origin string `json:"origin"`
}

func main() {
	// 初始化
	InitRedis()

	// 开启订阅并处理
	Sub()
}

// 加载配置, 创建REDIS客户端, 开启多协程处理
func InitRedis() {

	//
	config.ParseConfig()

	// 创建通道, 设置队列 [最大缓冲数量]
	queue = make(chan string, config.GlobalConfig.Queue)

	// REDIS 客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", config.GlobalConfig.RedisHost, config.GlobalConfig.RedisPort),
		Password:   config.GlobalConfig.RedisPass,
		MaxRetries: 2,
	})

	// 创建多协程, 同步处理
	for i := 0; i < config.GlobalConfig.Thread; i++ {
		go loop(i)
	}
}

// 开启REDIS 订阅 并处理
func Sub() {
	log.Printf("开启REDIS 订阅 并处理 ... ")
	// 订阅
	sub := redisClient.Subscribe(config.GlobalConfig.RedisChannel)

	count := 0
	for {
		// 接收消息 [阻塞等待]
		msg, err := sub.ReceiveMessage()
		if err != nil && err != io.EOF {
			log.Printf("获取数据出错 ==> [ %v ]", err)
			continue
		}

		// 发往通道队列
		if msg != nil {
			queue <- msg.Payload
			count++
		}
	}
}

// 多协程处理
func loop(index int) {
	for msg := range queue {
		// 处理队列消息
		handle(msg)
	}
}

// 解析消息
func handle(str string) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("解析消息失败 [ %s ] ==> [ %v ]", str, e)
		}
	}()

	// 转换成纯净JSON
	str = str[1 : len(str)-1]
	str = strings.Replace(str, "\\", "", -1)

	// 解析消息
	msg := &Message{}
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		log.Printf("解析json出错 ==> [ %v ] [ %v ]", str, err)
		return
	}

	// 打印消息
	log.Printf("订阅消息 ==> [ %v ] ", msg)
}
