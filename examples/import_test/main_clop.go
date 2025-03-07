// Code generated by clop-gen; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// TestImport的Parse方法，用于解析命令行参数
func (c *TestImport) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 设置默认值
	if ip := net.ParseIP("127.0.0.1"); ip != nil {
		c.IP = ip
	}

	if u, err := url.Parse("https://example.com"); err == nil {
		c.URL = *u
	}

	// 检查是否指定了配置文件
	var configFile string
	for i := 0; i < len(args); i++ {
		if args[i] == "--config" || args[i] == "-c" {
			if i+1 < len(args) {
				configFile = args[i+1]
				// 移除 --config 和它的值，避免后续解析错误
				if i+2 < len(args) {
					args = append(args[:i], args[i+2:]...)
				} else {
					args = args[:i]
				}
				i--
				break
			}
		}
	}

	// 从配置文件加载选项
	if configFile != "" {
		if err := c.loadFromConfigFile(configFile); err != nil {
			return fmt.Errorf("加载配置文件失败: %w", err)
		}
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--help" || arg == "-h":
			c.Usage()
			os.Exit(0)

		case arg == "--ip":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --ip")
			}

			if ip := net.ParseIP(args[i+1]); ip != nil {
				c.IP = ip
			}

			i++

		case arg == "--ip-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --ip-ptr")
			}

			if ip := net.ParseIP(args[i+1]); ip != nil {
				c.IPPtr = &ip
			}

			i++

		case arg == "--ipnet":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --ipnet")
			}

			if _, ipNet, err := net.ParseCIDR(args[i+1]); err == nil {
				c.IPNet = *ipNet
			}

			i++

		case arg == "--ipnet-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --ipnet-ptr")
			}

			if _, ipNet, err := net.ParseCIDR(args[i+1]); err == nil {
				c.IPNetPtr = ipNet
			}

			i++

		case arg == "--url":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --url")
			}

			if u, err := url.Parse(args[i+1]); err == nil {
				c.URL = *u
			}

			i++

		case arg == "--url-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --url-ptr")
			}

			if u, err := url.Parse(args[i+1]); err == nil {
				c.URLPtr = u
			}

			i++

		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

// UsageTmpl 是用于生成 Usage 函数的模板
func (c *TestImport) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    --ip    IP地址 (default: 127.0.0.1)
    --ip-ptr    IP地址指针
    --ipnet    IP网络
    --ipnet-ptr    IP网络指针
    --url    URL (default: https://example.com)
    --url-ptr    URL指针
`)
}

func (c *TestImport) loadFromConfigFile(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 根据文件扩展名选择解析方式
	ext := filepath.Ext(configFile)
	switch strings.ToLower(ext) {
	case ".json":
		return json.Unmarshal(data, c)
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, c)
	case ".toml":
		_, err := toml.Decode(string(data), c)
		return err
	default:
		return fmt.Errorf("不支持的配置文件格式: %s", ext)
	}
}
