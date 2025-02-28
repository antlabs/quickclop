package main

import (
	"fmt"
	"os"
)

// Config 是一个示例配置结构体，演示环境变量支持
// :quickclop
type Config struct {
	Debug    bool   `clop:"-d; --debug" usage:"启用调试模式" env:"APP_DEBUG"`
	Port     int    `clop:"-p; --port" usage:"服务端口" env:"APP_PORT"`
	Host     string `clop:"-h; --host" usage:"服务主机" env:"APP_HOST"`
	LogLevel string `clop:"--log-level" usage:"日志级别" env:"APP_LOG_LEVEL"`
	DataDir  string `clop:"args=data-dir" usage:"数据目录" env:"APP_DATA_DIR"`
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
	fmt.Printf("  DataDir:  %s\n", config.DataDir)
}
