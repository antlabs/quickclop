// Code generated by clop-gen; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Options的Parse方法，用于解析命令行参数
func (c *Options) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 设置默认值
	c.String = "default string"
	if val, err := strconv.Atoi("42"); err == nil {
		c.Int = val
	}
	if val, err := strconv.ParseFloat("3.14", 64); err == nil {
		c.Float = val
	}
	if duration, err := time.ParseDuration("1h30m"); err == nil {
		c.Duration = duration
	}

	defaultStr := "default ptr"
	c.StringPtr = &defaultStr
	if val, err := strconv.Atoi("100"); err == nil {
		c.IntPtr = &val
	}
	if val, err := strconv.ParseFloat("2.718", 64); err == nil {
		c.FloatPtr = &val
	}
	if val, err := strconv.ParseBool("true"); err == nil {
		c.BoolPtr = &val
	}

	c.StringSlice = strings.Split("a,b,c", ",")

	// 从环境变量中读取值

	if envVal := os.Getenv("QUICKCLOP_TEST_ENV"); envVal != "" {

		c.EnvValue = envVal

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

		case arg == "-s" || arg == "--string":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -s or --string")
			}

			c.String = args[i+1]

			i++

		case arg == "-i" || arg == "--int":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -i or --int")
			}

			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for -i or --int: %v", err)
			}
			c.Int = val

			i++

		case arg == "-f" || arg == "--float":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -f or --float")
			}

			val, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return fmt.Errorf("invalid value for -f or --float: %v", err)
			}
			c.Float = val

			i++

		case arg == "-b" || arg == "--bool":

			c.Bool = true

		case arg == "-d" || arg == "--duration":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -d or --duration")
			}

			if duration, err := time.ParseDuration(args[i+1]); err == nil {
				c.Duration = duration
			}

			i++

		case arg == "--string-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --string-ptr")
			}

			val := args[i+1]
			c.StringPtr = &val

			i++

		case arg == "--int-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --int-ptr")
			}

			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("invalid value for --int-ptr: %v", err)
			}
			c.IntPtr = &val

			i++

		case arg == "--float-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --float-ptr")
			}

			val, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return fmt.Errorf("invalid value for --float-ptr: %v", err)
			}
			c.FloatPtr = &val

			i++

		case arg == "--bool-ptr":

			val := true
			c.BoolPtr = &val

		case arg == "--strings":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --strings")
			}

			c.StringSlice = append(c.StringSlice, strings.Split(args[i+1], ",")...)

			i++

		case arg == "--env":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --env")
			}

			c.EnvValue = args[i+1]

			i++

		case arg == "-h" || arg == "--help":

			c.Help = true

		default:
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	return nil
}

// UsageTmpl 是用于生成 Usage 函数的模板
func (c *Options) Usage() {
	fmt.Println(`USAGE:
  myapp [OPTIONS]

OPTIONS:
    -c, --config    指定配置文件路径 (支持 JSON, YAML, TOML)
    -h, --help      显示帮助信息
    -s, --string    A string value (default: default string)
    -i, --int    An integer value (default: 42)
    -f, --float    A float value (default: 3.14)
    -b, --bool    A boolean flag
    -d, --duration    A duration value (default: 1h30m)
    --string-ptr    A string pointer value (default: default ptr)
    --int-ptr    An integer pointer value (default: 100)
    --float-ptr    A float pointer value (default: 2.718)
    --bool-ptr    A boolean pointer flag (default: true)
    --strings    A string slice (default: a,b,c)
    --env    Value from environment [env: QUICKCLOP_TEST_ENV]
    -h, --help    Show help information
`)
}

func (c *Options) loadFromConfigFile(configFile string) error {
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
