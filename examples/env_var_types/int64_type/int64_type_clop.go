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
)

// Options的Parse方法，用于解析命令行参数
func (c *Options) Parse(args []string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	// 从环境变量中读取值

	if envVal := os.Getenv("INT64_VALUE"); envVal != "" {

	}

	if envVal := os.Getenv("INT64_PTR_VALUE"); envVal != "" {

	}

	if envVal := os.Getenv("INT64_SLICE_VALUE"); envVal != "" {

		for _, s := range strings.Split(envVal, ",") {
			if val, err := strconv.ParseInt(s, 10, 64); err == nil {
				c.Int64Slice = append(c.Int64Slice, val)
			}
		}

	}

	if envVal := os.Getenv("INT64_SLICE_PTR_VALUE"); envVal != "" {

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

		case arg == "-i" || arg == "--int64":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for -i or --int64")
			}

			val, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid value for -i or --int64: %v", err)
			}
			c.Int64 = val

			i++

		case arg == "--int64-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --int64-ptr")
			}

			val, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid value for --int64-ptr: %v", err)
			}
			c.Int64Ptr = &val

			i++

		case arg == "--int64-slice":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --int64-slice")
			}

			for _, s := range strings.Split(args[i+1], ",") {
				if val, err := strconv.ParseInt(s, 10, 64); err == nil {
					c.Int64Slice = append(c.Int64Slice, val)
				}
			}

			i++

		case arg == "--int64-slice-ptr":

			if i+1 >= len(args) {
				return fmt.Errorf("missing value for --int64-slice-ptr")
			}

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
    -i, --int64    An int64 value [env: INT64_VALUE]
    --int64-ptr    An int64 pointer value [env: INT64_PTR_VALUE]
    --int64-slice    A slice of int64 values [env: INT64_SLICE_VALUE]
    --int64-slice-ptr    A pointer to a slice of int64 values [env: INT64_SLICE_PTR_VALUE]
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
