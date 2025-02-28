package main

import (
	"fmt"
	"os"
)

// Config 是一个示例配置结构体，演示默认值支持
// :quickclop
type Config struct {
	Debug    bool   `clop:"-d; --debug" usage:"启用调试模式" default:"false"`
	Port     int    `clop:"-p; --port" usage:"服务端口" default:"8080"`
	Host     string `clop:"-h; --host" usage:"服务主机" default:"localhost"`
	LogLevel string `clop:"--log-level" usage:"日志级别" default:"info"`
	Timeout  int    `clop:"-t; --timeout" usage:"超时时间（秒）" default:"30" env:"APP_TIMEOUT"`
	DataDir  string `clop:"args=data-dir" usage:"数据目录" default:"./data"`
}

func main() {
	config := &Config{}
	if err := config.Parse(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		config.Usage()
		os.Exit(1)
	}

	fmt.Println("配置信息:")
	fmt.Printf("  Debug:    %v\n", config.Debug)
	fmt.Printf("  Port:     %d\n", config.Port)
	fmt.Printf("  Host:     %s\n", config.Host)
	fmt.Printf("  LogLevel: %s\n", config.LogLevel)
	fmt.Printf("  Timeout:  %d\n", config.Timeout)
	fmt.Printf("  DataDir:  %s\n", config.DataDir)
}
