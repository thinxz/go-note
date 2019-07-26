package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	GlobalConfig = &Config{}
)

// ---------- ---------- ----------
// 配置文件配置项定义
// ---------- ---------- ----------
type Config struct {
	Dev          bool     `json:"dev"`
	Queue        int      `json:"queue"`
	Thread       int      `json:"thread"`
	RedisHost    string   `json:"redis_host"`
	RedisPort    int      `json:"redis_port"`
	RedisPass    string   `json:"redis_pass"`
	RedisChannel string   `json:"redis_channel"`
	Kafka        []string `json:"kafka"`
	KafkaTopic   string   `json:"kafka_topic"`
	Port         int      `json:"port"`
}

func main() {
	// 加载配置文件绝对路径
	configPath := flag.String("c", "c", "")
	flag.Parse()

	//
	log.Println(GetAbsolutePath())

	// 打开配置文件
	configFile, _ := os.Open(*configPath)
	configFileBytes, _ := ioutil.ReadAll(configFile)

	// JSON 格式解析成对象
	err := json.Unmarshal(configFileBytes, GlobalConfig)
	if err != nil {
		log.Panic("读取配置文件出错", err)
		return
	}

	// 输出配置
	log.Println(GlobalConfig)
}

// 获取运行时, 程序所在的绝对路径
func GetAbsolutePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
