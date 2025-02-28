package main

import (
	"fmt"
	"os"
	"time"
)

// Config 是一个示例配置结构体，演示时间类型支持
// :quickclop
type Config struct {
	StartTime  time.Time     `clop:"--start-time" usage:"开始时间" default:"2025-01-01T00:00:00Z"`
	Duration   time.Duration `clop:"-d; --duration" usage:"持续时间" default:"24h"`
	Timeout    int           `clop:"-t; --timeout" usage:"超时时间（秒）" default:"30"`
	EnableLogs bool          `clop:"--enable-logs" usage:"启用日志" default:"true"`
}

func main() {
	config := &Config{}
	if err := config.Parse(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		config.Usage()
		os.Exit(1)
	}

	fmt.Println("配置信息:")
	fmt.Printf("  StartTime: %v\n", config.StartTime)
	fmt.Printf("  Duration:  %v\n", config.Duration)
	fmt.Printf("  Timeout:   %d\n", config.Timeout)
	fmt.Printf("  EnableLogs: %v\n", config.EnableLogs)
}
